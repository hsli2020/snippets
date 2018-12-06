<?php

/* connect to gmail */
$hostname = '{imap.gmail.com:993/imap/ssl}INBOX';
$username = 'davidwalshblog@gmail.com';
$password = 'davidwalsh';

/* try to connect */
$inbox = imap_open($hostname,$username,$password) or die('Cannot connect to Gmail: ' . imap_last_error());

/* grab emails */
$emails = imap_search($inbox,'ALL');

/* if emails are returned, cycle through each... */
if($emails) {

	/* begin output var */
	$output = '';

	/* put the newest emails on top */
	rsort($emails);

	/* for every email... */
	foreach($emails as $email_number) {

		/* get information specific to this email */
		$overview = imap_fetch_overview($inbox,$email_number,0);
		$message = imap_fetchbody($inbox,$email_number,2);

		/* output the email header information */
		$output.= '<div class="toggler '.($overview[0]->seen ? 'read' : 'unread').'">';
		$output.= '<span class="subject">'.$overview[0]->subject.'</span> ';
		$output.= '<span class="from">'.$overview[0]->from.'</span>';
		$output.= '<span class="date">on '.$overview[0]->date.'</span>';
		$output.= '</div>';

		/* output the email body */
		$output.= '<div class="body">'.$message.'</div>';
	}

	echo $output;
}

/* close the connection */
imap_close($inbox);

/*

for ($i = 1; $i <= $count; $i++) {
    $header = imap_headerinfo($connection, $i);
    $raw_body = imap_body($connection, $i);
}

Header object

The value returned from the imap_headerinfo() function is an object containing
a number of properties. To show the values available, an example of the output
from print_r is as follows:

stdClass Object
(
    [date] => Wed, 4 Feb 2009 22:37:42 +1300
    [Date] => Wed, 4 Feb 2009 22:37:42 +1300
    [subject] => Fwd: another test
    [Subject] => Fwd: another test
    [in_reply_to] => <59f89e00902040137vb73ed1ep5f870dafe02f26cf@mail.example.com>
    [message_id] => <59f89e00902040137s16c317b6oa4658a4d2cc64c3c@mail.example.com>
    [references] => <59f89e00902040137vb73ed1ep5f870dafe02f26cf@mail.example.com>
    [toaddress] => john@example.com
    [to] => Array
    (
        [0] => stdClass Object
        (
            [mailbox] => john
            [host] => example.com
        )
    )

    [fromaddress] => Chris Hope
    [from] => Array
    (
        [0] => stdClass Object
        (
            [personal] => Chris Hope
            [mailbox] => chris
            [host] => example.com
        )
    )

    [reply_toaddress] => Chris Hope
    [reply_to] => Array
    (
        [0] => stdClass Object
        (
            [personal] => Chris Hope
            [mailbox] => chris
            [host] => example.com
        )
    )

    [senderaddress] => Chris Hope
    [sender] => Array
    (
        [0] => stdClass Object
        (
            [personal] => Chris Hope
            [mailbox] => chris
            [host] => example.com
        )
    )

    [Recent] => N
    [Unseen] =>
    [Flagged] =>
    [Answered] =>
    [Deleted] =>
    [Draft] =>
    [Msgno] =>   20
    [MailDate] =>  4-Feb-2009 22:37:42 +1300
    [Size] => 3111
    [udate] => 1233740262
)


