<?php

/**
 * Send a Message to a Slack Channel.
 *
 * In order to get the API Token visit: https://api.slack.com/custom-integrations/legacy-tokens
 * The token will look something like this `xoxo-2100000415-0000000000-0000000000-ab1ab1`.
 * 
 * @param string $message The message to post into a channel.
 * @param string $channel The name of the channel prefixed with #, example #foobar
 * @return boolean
 */
function slack($message, $channel)
{
    $ch = curl_init("https://slack.com/api/chat.postMessage");
    $data = http_build_query([
        "token"    => "<YOUR-SLACK-API-TOKEN>",
    	"channel"  => $channel, //"#mychannel",
    	"text"     => $message, //"Hello, Foo-Bar channel message.",
    	"username" => "MySlackBot",
    ]);
    curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'POST');
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
    $result = curl_exec($ch);
    curl_close($ch);
    
    return $result;
}

slack('Amazon order 111-2222222-3333333 cancelled', '#general');
