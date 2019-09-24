<?php

function getTariffCode($desc,$uprice)
{
	$default = "8471.30.0100";  // adaptor

    $map = [
        "projector lamp" => "8528.62.0000",
        "projector lens" => "8528.62.0000",
        "label printer"  => "8471.60.90.50",
        "photo printer"  => "8471.30.0100",
        "scanner"        => "8471.60.9050",
        "monitor"        => "8528.52.0000",
        "printer"        => "8471.60.9050",
        "projector"      => "8528.62.0000",
        "power supply"   => "8504.40.6018",
        "laptop"         => "8471.30.0100",
        "notebook"       => "8471.30.0100",
        "vivobook"       => "8471.30.0100",
        "graphics card"  => "8471.30.0100",
        "keyboard"       => "8471.60.2000",
        "desktop"        => "8471.41.0150",
        "ssd"            => "8523.51.0000",
        "kitchen"        => "8208.30.00",
        "irobot"         => "8208.30.00",
        "motherboard"    => "8471.30.0100",
        "access point"   => "8517.62.90",
        "monitor"        => "8528.52.0000",
        "software"       => "8523.80.20",
        "adapter"        => "8471.30.0100",
        "camera"         => "8525.80.4000",
        "cable"          => "8471.30.0100",
        "mini pc"        => "8471.30.0100",
        "paper"          => "4811.59.2000",
        "sewing machine" => "8452.10.00",
        "video record"   => "8528.49.0500",
        "tablet"         => "8471.30.0100",
        "tower case"     => "8473.30.5100",
        "cartridge"      => "8443.99.20.10",
        "smartphone"     => "8517.12.0500",
        "speaker"        => "8518.29.8000",
        "shredder"       => "8441.30.0100",
        "router"         => "8471.30.0100",
        "roller set"     => "8471.30.0100",
        "phone"          => "8511.11.00.00",
        "network"        => "8473.30.1180",
        "mouse"          => "8471.60.2000",
        "mount"          => "8473.30.5100",
        "microphone"     => "8518.10.40.00",
        "lock"           => "8473.30.5100",
        "dvd player"     => "8519.81.10",
        "education"      => "8543.70.9301",
        "wheeled toys"   => "9503.00.00",
        "game console"   => "8526.92.1000",
        "controller"     => "8526.92.1000",
        "board"          => "4410.11.00",
        "calculator"     => "8473.21.00.00",
        "backpack"       => "8471.30.10.00",
        "amplifier"      => "8542.33.00",
        "power strip"    => "8504.40.6018",
        "scan"           => "8471.60.9050",
        "workstation"    => "8471.41.0150",
    ];

    foreach ($map as $name =>$code) {
        if (strstr($desc, $name) != false) {
            return $code;
        }
    }

    return $default;
}

function isFragile($desc)
{
    $list = [ "monitor","printer","scanner" ];

    foreach ($list as $kwd) {
        if (strstr($desc, $kwd) != false) {
            return 1;
        }
    }

    return 0;
}

