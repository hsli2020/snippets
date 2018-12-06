<?php

// obviously change these to your actual login details and server connection string
$login = 'me@example.com';
$password = 'mypassword';
$server = '{imap.gmail.com:993/ssl/novalidate-cert}';

// connect to server
$connection = imap_open($server, $login, $password);

// download email message #1 and fetch it
$emailMessage = new EmailMessage($connection, 1);
$emailMessage->fetch();

// match inline images in html content
preg_match_all('/src="cid:(.*)"/Uims', $emailMessage->bodyHTML, $matches);

// if there are any matches, loop through them and save to filesystem, change the src property
// of the image to an actual URL it can be viewed at
if(count($matches)) {

	// search and replace arrays will be used in str_replace function below
	$search = array();
	$replace = array();

	foreach($matches[1] as $match) {
		// work out some unique filename for it and save to filesystem etc
		$uniqueFilename = "UNIQUE_FILENAME.extension";
		// change /path/to/images to actual path
		file_put_contents("/path/to/images/$uniqueFilename", $emailMessage->attachments[$match]['data']);
		$search[] = "src=\"cid:$match\"";
		// change www.example.com etc to actual URL
		$replace[] = "src=\"http://www.example.com/images/$uniqueFilename\"";
	}

	// now do the replacements
	$emailMessage->bodyHTML = str_replace($search, $replace, $emailMessage->bodyHTML);
}

// all done, close the connection to the mail server
imap_close($connection);
