<?php
#!/usr/bin/php -q
/*
-- Here's my DB structure

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- Table structure for table `emails`

CREATE TABLE IF NOT EXISTS `emails` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `from` varchar(250) NOT NULL,
  `subject` text NOT NULL,
  `body` text NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1;

-- --------------------------------------------------------
-- Table structure for table `files`

CREATE TABLE IF NOT EXISTS `files` (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `email_id` int(100) NOT NULL,
  `filename` varchar(255) NOT NULL,
  `size` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1;


CREATE TABLE `mails` (
  `message_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `message` varchar(10000) NOT NULL DEFAULT '',
  `file` longblob,
  `mailingdate` varchar(40) DEFAULT NULL,
  `starred_status` int(10) unsigned NOT NULL DEFAULT '0',
  `sender_email` varchar(200) NOT NULL DEFAULT '',
  `reciever_email` varchar(200) NOT NULL DEFAULT '',
  `inbox_status` int(10) unsigned NOT NULL DEFAULT '0',
  `sent_status` int(10) unsigned NOT NULL DEFAULT '0',
  `draft_status` int(10) unsigned NOT NULL DEFAULT '0',
  `trash_status` int(10) unsigned NOT NULL DEFAULT '0',
  `subject` varchar(200) DEFAULT NULL,
  `read_status` int(10) unsigned NOT NULL DEFAULT '0',
  `delete_status` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`message_id`)
)
*/

//  Use -q so that php doesn't print out the HTTP headers
//  Anything printed to STDOUT will be sent back to the sender as an error!

//  Config options

$max_time_limit = 600; // in seconds
// A safe place for files with trailing slash (malicious users could upload a php or executable file!)
$save_directory = "/some/folder/path";
$allowedSenders = Array('myemail@gmail.com',
    'Bob the Builder <whatever@whoever.com>'); // only people you trust!

$send_email = TRUE; // Send confirmation e-mail?
$save_msg_to_db = TRUE; // Save e-mail body to DB?

$db_host = 'localhost';
$db_un = 'db_un';
$db_pass = 'password';
$db_name = 'db_name';

// ------------------------------------------------------

set_time_limit($max_time_limit);
ini_set('max_execution_time',$max_time_limit);

global $from, $subject, $boundary, $message, $save_path,$files_uploaded;
$save_path = $save_directory;
$files_uploaded = Array();

function formatBytes(&$bytes, $precision = 2) {
    $units = array('B', 'KB', 'MB', 'GB', 'TB');

    $bytes = max($bytes, 0);
    $pow = floor(($bytes ? log($bytes) : 0) / log(1024));
    $pow = min($pow, count($units) - 1);

    $bytes /= pow(1024, $pow);

    return round($bytes, $precision) . ' ' . $units[$pow];
}

function process_part(&$email_part){
    global $message;

    // Max two parts. The data could have more than one \n\n in it somewhere,
    // but the first \n\n should be after the content info block
    $parts = explode("\n\n",$email_part,2);

    $info = split("\n",$parts[0]);
    $type;
    $name;
    $encoding;
    foreach($info as $line){
        if(preg_match("/Content-Type: (.*);/",$line,$matches)){
            $type = $matches[1];
        }
        if(preg_match("/Content-Disposition: attachment; filename=\"(.*)\"/",
            $line,$matches)){
            $name = time() . "_" . $matches[1];
        }
        if(preg_match("/Content-Transfer-Encoding: (.*)/",$line,$matches)){
            $encoding = $matches[1];
        }
    }

    // We don't know what it is, so we don't know how to process it
    if(!isset($type)){ return FALSE; }

    switch($type){
    case 'text/plain':
        // "But if you get a text attachment, you're going to overwrite
        // the real message!" Yes. I don't care in this case...
        $message = $parts[1];
        break;
    case 'multipart/alternative':
        // Multipart comes where the client sends the data in two formats so
        // that recipients who can't read (or don't like) fancy content
        // have another way to read it. Eg. When sending an html formatted
        // message, they will also send a plain text message
        process_multipart($info,$parts[1]);
        break;
    default:
        if(isset($name)){ // the main message will not have a file name...
            // text/html messages won't be saved!
            process_data($name,$encoding,$parts[1]);
        }elseif(!isset($message) && strpos($type,'text') !== FALSE){
            $message = $parts[1]; // Going out on a limb here...capture
            // the message
        }
        break;
    }
}

function process_multipart(&$info,&$data){
    global $message;

    $bounds;
    foreach($info as $line){
        if (preg_match("/boundary=(.*)$/",$line,$matches)){
            $bounds = $matches[1];
        }
    }

    $multi_parts = split("--" .$bounds,$data);
    for($i = 1;$i < count($multi_parts);$i++){
        process_part($multi_parts[$i]);
    }
}

function process_data(&$name,&$encoding = 'base64' ,&$data){
    global $save_path,$files_uploaded;

    // find a filename that's not in use. There's a race condition
    // here which should be handled with flock or something instead
    // of just checking for a free filename

    $unlocked_and_unique = FALSE;
    while(!$unlocked_and_unique){
        // Find unique
        $name = time() . "_" . $name;
        while(file_exists($save_path . $name)) {
            $name = time() . "_" . $name;
        }

        // Attempt to lock
        $outfile = fopen($save_path.$name,'w');
        if(flock($outfile,LOCK_EX)){
            $unlocked_and_unique = TRUE;
        }else{
            flock($outfile,LOCK_UN);
            fclose($outfile);
        }
    }

    if($encoding == 'base64'){
        fwrite($outfile,base64_decode($data));
    }elseif($encoding == 'uuencode'){
        // I haven't actually seen this in an e-mail, but older clients may
        // still use it...not 100% sure that this will work correctly as is
        fwrite($outfile,convert_uudecode($data));
    }
    flock($outfile,LOCK_UN);
    fclose($outfile);

    // This is for readability for the return e-mail and in the DB
    $files_uploaded[$name] = formatBytes(filesize($save_path.$name));
}

// Process the e-mail from stdin
$fd = fopen('php://stdin','r');
$email = '';
while(!feof($fd)){ $email .= fread($fd,1024); }

// Headers hsould go till the first \n\n. Grab everything before that, then
// split on \n and process each line
$headers = split("\n",array_shift(explode("\n\n",$email,2)));
foreach($headers as $line){
    // The only 3 headers we care about...
    if (preg_match("/^Subject: (.*)/", $line, $matches)) {
        $subject = $matches[1];
    }
    if (preg_match("/^From: (.*)/", $line, $matches)) {
        $from = $matches[1];
    }
    if (preg_match("/boundary=(.*)$/",$line,$matches)){
        $boundary = $matches[1];
    }
}

// Check $from here and exit if it's blank or
// not someone you want to get mail from!
if(!in_array($from,$allowedSenders)){
    die("Not an allowed sender");
}

// No boundary was in the e-mail sent to us. We don't know what to do!
if(!isset($boundary)){
    die("I couldn't find an e-mail boundary. Maybe this isn't an e-mail");
}

// Split the e-mail on the found boundary
// The first part will be the header (hence $i = 1 in our loop)
// Each other chunk should have some info on the chunk,
// then \n\n then the chunk data
// Process each chunk
$email_parts = split("--" . $boundary,$email);
for($i = 1;$i < count($email_parts);$i++){
    process_part($email_parts[$i]);
}

// Put the results in the database if needed
if($save_msg_to_db){
    mysql_connect($db_host,$db_un,$db_pass);
    mysql_select_db($db_name);

    $q = "INSERT INTO `emails` (`from`,`subject`,`body`) VALUES ('" .
	mysql_real_escape_string($from) . "','" .
	mysql_real_escape_string($subject) . "','" .
	mysql_real_escape_string($message) . "')";

    mysql_query($q) or die(mysql_error());

    if(count($files_uploaded) > 0){
        $id = mysql_insert_id();
        $q = "INSERT INTO `files` (`email_id`,`filename`,`size`) VALUES ";
        $filesar = Array();
        foreach($files_uploaded as $f => $s){
            $filesar[] = "('$id','" .
                mysql_real_escape_string($f) . "','" .
                mysql_real_escape_string($s) . "')";
        }
        $q .= implode(', ',$filesar);
        mysql_query($q) or die(mysql_error());
    }
}

// Send response e-mail if needed
if($send_email && $from != ""){
    $to = $from;
    $newmsg = "Thanks! I just uploaded the following ";
    $newmsg .= "files to your storage:\n\n";
    $newmsg .= "Filename -- Size\n";
    foreach($files_uploaded as $f => $s){
        $newmsg .= "$f -- $s\n";
    }
    $newmsg .= "\nI hope everything looks right. If not,";
    $newmsg .=  "please send me an e-mail!\n";

    mail($to,$subject,$newmsg);
}
