<?php

function pop3_login($host,$port,$user,$pass,$folder="INBOX",$ssl=false)
{
    $ssl=($ssl==false)?"/novalidate-cert":"";
    return (imap_open("{"."$host:$port/pop3$ssl"."}$folder",$user,$pass));
}

function pop3_stat($connection)
{
    $check = imap_mailboxmsginfo($imap);
    return ((array)$check);
}

function pop3_list($connection,$message="")
{
    if ($message)
    {
        $range=$message;
    } else {
        $MC = imap_check($mbox);
        $range = "1:".$MC->Nmsgs;
    }
    $response = imap_fetch_overview($mbox,$range);
    foreach ($response as $msg) $result[$msg->msgno]=(array)$msg;
}

function pop3_retr($connection,$message)
{
    return(imap_fetchheader($connection,$message,FT_PREFETCHTEXT));
}

function pop3_dele($connection,$message)
{
    return(imap_delete($connection,$message));
}

function mail_parse_headers($headers)
{
    $headers=preg_replace('/\r\n\s+/m', '',$headers);
    preg_match_all('/([^: ]+): (.+?(?:\r\n\s(?:.+?))*)?\r\n/m', $headers, $matches);
    foreach ($matches[1] as $key =>$value) $result[$value]=$matches[2][$key];
    return($result);
}

function mail_mime_to_array($imap,$mid,$parse_headers=false)
{
    $mail = imap_fetchstructure($imap,$mid);
    $mail = mail_get_parts($imap,$mid,$mail,0);
    if ($parse_headers) $mail[0]["parsed"]=mail_parse_headers($mail[0]["data"]);
    return($mail);
}

function mail_get_parts($imap,$mid,$part,$prefix)
{
    $attachments=array();
    $attachments[$prefix]=mail_decode_part($imap,$mid,$part,$prefix);
    if (isset($part->parts)) // multipart
    {
        $prefix = ($prefix == "0")?"":"$prefix.";
        foreach ($part->parts as $number=>$subpart)
            $attachments=array_merge($attachments, mail_get_parts($imap,$mid,$subpart,$prefix.($number+1)));
    }
    return $attachments;
}

function mail_decode_part($connection,$message_number,$part,$prefix)
{
    $attachment = array();

    if($part->ifdparameters) {
        foreach($part->dparameters as $object) {
            $attachment[strtolower($object->attribute)]=$object->value;
            if(strtolower($object->attribute) == 'filename') {
                $attachment['is_attachment'] = true;
                $attachment['filename'] = $object->value;
            }
        }
    }

    if($part->ifparameters) {
        foreach($part->parameters as $object) {
            $attachment[strtolower($object->attribute)]=$object->value;
            if(strtolower($object->attribute) == 'name') {
                $attachment['is_attachment'] = true;
                $attachment['name'] = $object->value;
            }
        }
    }

    $attachment['data'] = imap_fetchbody($connection, $message_number, $prefix);
    if($part->encoding == 3) { // 3 = BASE64
        $attachment['data'] = base64_decode($attachment['data']);
    }
    elseif($part->encoding == 4) { // 4 = QUOTED-PRINTABLE
        $attachment['data'] = quoted_printable_decode($attachment['data']);
    }
    return($attachment);
}
