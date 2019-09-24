<?php

if (count($argv) != 2) {
    echo "usage php xmltogo.php filename.xml\n";
    exit;
}

$text = file_get_contents($argv[1]);
$xml = simplexml_load_string($text);

echo "func Test_Response(t *testing.T) {\n";
echo "\tassert := assert.New(t)\n\n";

echo "\tfilename := \"fixtures/Response.xml\"\n";
echo "\tdata := getXML(filename)\n\n";

echo "\tvar x ", makeVarname($xml->getName()), "\n\n";

echo "\terr := xml.Unmarshal(data, &x)\n";
echo "\tassert.Nil(err)\n\n";

walkxml($xml, 'x');

echo "}\n";

function walkxml($xml, $label)
{
    $xmltag = $xml->getName();
    $varname = makeVarname($xmltag);

	if (count($xml->children()) == 0) {
		codeln("$label.$varname", strval($xml));
	}

    foreach($xml->attributes() as $name => $val) {
        $name = makeVarname($name);
        codeln("$label.$varname.$name", $val);
    }

    foreach ($xml as $node) {
        $xmltag = $node->getName();
        $name = makeVarname($xmltag);

        if ($node->children()) {
            walkxml($node, "$label.$varname");
        } else {
            codeln("$label.$varname.$name", strval($node));
        }
    }
}

function makeVarname($name)
{
    $arr = array_map('ucfirst', explode('-', $name));
    return implode('', $arr);
}

function codeln($name, $value)
{
    echo "\tassert.Equal(", $name, ', "', $value, '")', "\n";
}
