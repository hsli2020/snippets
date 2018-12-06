<?php
function sendMail($data) {
    $recipient = 'chuan.ma@avidlifemedia.com,carlos.nakhle@avidlifemedia.com';

    $csv = "pnum,msg_count,average_similar_text_score\n";

    $yesterday = date('Y-m-d', strtotime('yesterday'));
    $filename = "mail.php";

    // send the email
    $semi_rand = md5(time());
    $mime_boundary = "==Multipart_Boundary_x{$semi_rand}x";
    $header = "MIME-Version: 1.0\n" .
              "Content-Type: multipart/mixed;\n" .
              " boundary=\"{$mime_boundary}\"";

    $subject = "Female SPAM suspect report on $yesterday.";

    $body = "Hi, there\n\n"
          . "This mail is sent to you by FemalSpamSuspectList service of ashleymadison.com.\n"
          . "Please see attachment for the female SPAM suspect report of yesterday.\n\n"
          . "In the report,\n"
          . "1) average_similar_text_score is between 0 and 100.\n"
          . "2) bigger score means the text messages are more similar, 100 means the messages are completely the same.\n"
          . "3) only 11 random messages are picked each time for the similarity-msg-check.\n";

    $message = "This is a multi-part message in MIME format.\n\n" .
               "--{$mime_boundary}\n" .
               "Content-Type: text/plain; charset=\"iso-8859-1\"\n" .
               "Content-Transfer-Encoding: 7bit\n\n" .
               $body. "\n\n" .
               "--{$mime_boundary}\n" .
               "Content-Type: application/vnd.ms-excel;\n" .
               "Content-Disposition: attachment; filename=\"$filename\"\n" .
               "Content-Transfer-Encoding: base64\n\n" .
               chunk_split(base64_encode($csv)) . "\n\n" .
               "--{$mime_boundary}--\n";

    mail($recipient, $subject, $message, $header);
}

sendMail('');
