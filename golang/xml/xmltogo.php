<?php

if (count($argv) != 2) {
    echo "usage php xmltogo.php filename.xml\n";
    exit;
}

$text = file_get_contents($argv[1]);
$xml = simplexml_load_string($text);
walkxml($xml, 0);

function walkxml($xml, $indent)
{
    $xmltag = $xml->getName();
    $varname = makeVarname($xmltag);

    if ($xml->children()) {
        codeln($indent, "");
        codeln($indent, "type $varname struct {");
        $indent++;
        codeln($indent, "XMLName xml.Name ". '`xml:"'. $xmltag. '"`');
        foreach($xml->attributes() as $name => $_) {
            $varname = makeVarname($name);
            codeln($indent, $varname. ' string `xml:"'. $name. ',attr"`');
        }
        if (strlen(trim(strval($xml))) > 0) {
            codeln($indent, '//Value string `xml:",chardata"`');
        }
    } else {
        codeln($indent, "$varname string ". '`xml:"'. $xmltag. '"`');
    }

    foreach ($xml as $node) {
        $xmltag = $node->getName();
        $varname = makeVarname($xmltag);

        if ($node->children()) {
            walkxml($node, $indent);
        } else {
            codeln($indent, "$varname string ". '`xml:"'. $xmltag. '"`');
        }
    }

    if ($xml->children()) {
        $indent--;
        codeln($indent, "}\n");
    }
}

function makeVarname($name)
{
    $arr = array_map('ucfirst', explode('-', $name));
    return implode('', $arr);
}

function codeln($indent, $code)
{
    $spaces = str_repeat(' ', $indent*4);
    echo $spaces, $code, PHP_EOL;
}
