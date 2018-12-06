<?php
$note=<<<XML
<company>
  <staffs>
      <staff>
          <id>103</id>
          <firstname>Adam</firstname>
          <lastname>Ng</lastname>
          <username>adamng</username>
      </staff>
      <staff>
          <id>108</id>
          <firstname>Jennifer</firstname>
          <lastname>Loh</lastname>
          <username>jenniferloh</username>
      </staff>
  </staffs>
</company>
XML;

$xml=simplexml_load_string($note);
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
