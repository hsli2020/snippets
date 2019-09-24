<?php

query_v1();
query_v2();

function query_v2()
{
	$doc = new DOMDocument();
	$html = file_get_contents("PriceAlerts.html");
	@$doc->loadHTML($html);

	$xpath = new DOMXPath($doc);

    $cnt = 0;
	$tables = $doc->getElementsByTagName('table');
	foreach ($tables as $table) {
		$query = "tbody/tr/td/div/a";

		$entries = $xpath->query($query, $table);

		foreach ($entries as $entry) {
			$sku = trim($entry->nodeValue);
			//echo $sku, "\n";
            $cnt++;
		}
	}
    echo $cnt, "\n";
}

function query_v1()
{
	$doc = new DOMDocument();

	$html = file_get_contents("PriceAlerts.html");
	@$doc->loadHTML($html);

	$xpath = new DOMXPath($doc);

	$table = $doc->getElementsByTagName('table')->item(1);

	$query = "//tbody/tr/td/div/a";

	$entries = $xpath->query($query, $table);

    $cnt = 0;
	foreach ($entries as $entry) {
		$sku = trim($entry->nodeValue);
		//echo $sku, "\n";
        $cnt++;
	}
    echo $cnt, "\n";
}
