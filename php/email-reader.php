<?php

class Email_reader
{
	// imap server connection
	public $conn;

	// inbox storage and inbox message count
	private $inbox;
	private $msg_cnt;

	// email login credentials
	private $server = 'yourserver.com';
	private $user   = 'email@yourserver.com';
	private $pass   = 'yourpassword';
	private $port   = 143; // adjust according to server settings

	// connect to the server and get the inbox emails
    function __construct()
    {
		$this->connect();
		$this->inbox();
	}

	// close the server connection
    function close()
    {
		$this->inbox = array();
		$this->msg_cnt = 0;

		imap_close($this->conn);
	}

	// open the server connection
	// the imap_open function parameters will need to be changed for the particular server
	// these are laid out to connect to a Dreamhost IMAP server
    function connect()
    {
		$this->conn = imap_open('{'.$this->server.'/notls}', $this->user, $this->pass);
	}

	// move the message to a new folder
    function move($msg_index, $folder='INBOX.Processed')
    {
		// move on server
		imap_mail_move($this->conn, $msg_index, $folder);
		imap_expunge($this->conn);

		// re-read the inbox
		$this->inbox();
	}

	// get a specific message (1 = first email, 2 = second email, etc.)
    function get($msg_index=NULL)
    {
		if (count($this->inbox) <= 0) {
			return array();
		}
		elseif (!is_null($msg_index) && isset($this->inbox[$msg_index])) {
			return $this->inbox[$msg_index];
		}

		return $this->inbox[0];
	}

	// read the inbox
    function inbox()
    {
		$this->msg_cnt = imap_num_msg($this->conn);

		$in = array();
		for($i = 1; $i <= $this->msg_cnt; $i++) {
			$in[] = array(
				'index'     => $i,
				'header'    => imap_headerinfo($this->conn, $i),
				'body'      => imap_body($this->conn, $i),
				'structure' => imap_fetchstructure($this->conn, $i)
			);
		}

		$this->inbox = $in;
	}
}

function email_pull()
{
	// load the Email_reader library from previous post
	$this->load->library('email_reader');

	// load the meals_model to store meal information
	$this->load->model('meals_model');

	// this method is run on a cronjob and should process all emails in the inbox
	while (1) {
		// get an email
		$email = $this->email_reader->get();

		// if there are no emails, jump out
		if (count($email) <= 0) {
			break;
		}

		$attachments = array();
		// check for attachments
		if (isset($email['structure']->parts) && count($email['structure']->parts)) {
			// loop through all attachments
			for ($i = 0; $i < count($email['structure']->parts); $i++) {
				// set up an empty attachment
				$attachments[$i] = array(
					'is_attachment' => FALSE,
					'filename'      => '',
					'name'          => '',
					'attachment'    => ''
				);

				// if this attachment has idfparameters, then proceed
				if ($email['structure']->parts[$i]->ifdparameters) {
					foreach ($email['structure']->parts[$i]->dparameters as $object) {
						// if this attachment is a file, mark the attachment and filename
						if (strtolower($object->attribute) == 'filename') {
							$attachments[$i]['is_attachment'] = TRUE;
							$attachments[$i]['filename']      = $object->value;
						}
					}
				}

				// if this attachment has ifparameters, then proceed as above
				if ($email['structure']->parts[$i]->ifparameters) {
					foreach ($email['structure']->parts[$i]->parameters as $object) {
						if (strtolower($object->attribute) == 'name') {
							$attachments[$i]['is_attachment'] = TRUE;
							$attachments[$i]['name']          = $object->value;
						}
					}
				}

				// if we found a valid attachment for this 'part' of the email, process the attachment
				if ($attachments[$i]['is_attachment']) {
					// get the content of the attachment
					$attachments[$i]['attachment'] = imap_fetchbody($this->email_reader->conn, $email['index'], $i+1);

					// check if this is base64 encoding
					if ($email['structure']->parts[$i]->encoding == 3) { // 3 = BASE64
						$attachments[$i]['attachment'] = base64_decode($attachments[$i]['attachment']);
					}
					// otherwise, check if this is "quoted-printable" format
					elseif ($email['structure']->parts[$i]->encoding == 4) { // 4 = QUOTED-PRINTABLE
						$attachments[$i]['attachment'] = quoted_printable_decode($attachments[$i]['attachment']);
					}
				}
			}
		}

		// for My Slow Low, check if I found an image attachment
		$found_img = FALSE;
		foreach ($attachments as $a) {
			if ($a['is_attachment'] == 1) {
				// get information on the file
				$finfo = pathinfo($a['filename']);

				// check if the file is a jpg, png, or gif
				if (preg_match('/(jpg|gif|png)/i', $finfo['extension'], $n)) {
					$found_img = TRUE;
					// process the image (save, resize, crop, etc.)
					$fname = $this->_process_img($a['attachment'], $n[1]);

					break;
				}
			}
		}

		// if there was no image, move the email to the Rejected folder on the server
		if ( ! $found_img) {
			$this->email_reader->move($email['index'], 'INBOX.Rejected');
			continue;
		}

		// get content from the email that I want to store
		$addr   = $email['header']->from[0]->mailbox."@".$email['header']->from[0]->host;
		$sender = $email['header']->from[0]->mailbox;
		$text   = ( ! empty($email['header']->subject) ? $email['header']->subject : '');

		// move the email to Processed folder on the server
		$this->email_reader->move($email['index'], 'INBOX.Processed');

		// add the data to the database
		$this->meals_model->add(array(
			'username'    => $sender,
			'email'       => $addr,
			'photo'       => $fname,
			'description' => ($text == '' ? NULL : $text)
		));

		// don't slam the server
		sleep(1);
	}

	// close the connection to the IMAP server
	$this->email_reader->close();
}
