<?php

$text = 'Here is an example for a text (string) that contains one or more url. Just visit http://www.google.com/ or <a href="http://www.google.com" rel="nofollow">http://google.com</a> and this is the end of the example.';
$text = 'Here is an example for a text (string) that contains one or more url. Just visit <a href="http://www.google.com/">http://www.google.com/</a> or <a href="http://google.com">http://google.com</a> and this is the end of the example.';

function textWithLinks($str)
{
    return preg_replace(
        '@(https?://([-\w\.]+[-\w])+(:\d+)?(/([\w/_\.#-]*(\?\S+)?[^\.\s])?)?)@',
        '<a href="$1" target="_blank" rel="nofollow">$1</a>',
        strip_tags($str)
    );
}

$text = textWithLinks($text);
echo $text;

#-----------------------------------------------------------
# This is better

function makeUrltoLink($string)
{
    // The Regular Expression filter
    $pattern = "/(((http|https|ftp|ftps)\:\/\/)|(www\.))[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\:[0-9]+)?(\/\S*)?/";
     
    // make the urls to hyperlinks
    return preg_replace($pattern,
        '<a href="$0" target="_blank" rel="noopener noreferrer">$0</a>',
        strip_tags($string)
    );
}
 
$str = "Visit www.cluemediator.com and subscribe us on https://www.cluemediator.com/subscribe for regular updates.";
$str = 'Here is an example for a text (string) that contains one or more url. Just visit <a href="http://www.google.com/">http://www.google.com/</a> or <a href="http://google.com">http://google.com</a> and this is the end of the example.';
 
echo "\n\n";
echo $convertedStr = makeUrltoLink($str);
