<?php

class Base62{
	
	private static $base62 = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
	
    public static function encode($number, $encode = '')
    {
		while($number > 0){
			$mod = bcmod($number, 62);
			$encode .= self::$base62[$mod];
			$number = bcdiv(bcsub($number, $mod), 62);
			
		}
		return strrev($encode);
	}
	
    public static function decode($encode, $number = 0)
    {
		$length = strlen($encode);
		$baselist = array_flip(str_split(self::$base62));
		for($i = 0; $i < $length; $i++){
			$number = bcadd($number, bcmul($baselist[$encode[$i]],  bcpow(62, $length - $i - 1)));
		}	
		return $number;		
	}
}

//$str = Base62::encode(3696);
//$num = Base62::decode($str);
//echo $num, ' => ', $str, "\n";

//$str = Base62::encode(3697);
//$num = Base62::decode($str);
//echo $num, ' => ', $str, "\n";

$chars=array("a","b","c","d","e","f","g","h",
		     "i","j","k","l","m","n","o","p",
			 "q","r","s","t","u","v","w","x",
			 "y","z","0","1","2","3","4","5",
			 "6","7","8","9","A","B","C","D",
			 "E","F","G","H","I","J","K","L",
			 "M","N","O","P","Q","R","S","T",
			 "U","V","W","X","Y","Z");

$salt="www.joneto.com";
$hash=md5("http://www.sina.com".$salt);
$rs=array();
for($i=0; $i<4; $i++) {
	$temp=substr($hash, $i*8,8);
	$temp=base_convert($temp, 16, 10) & base_convert("3fffffff", 16, 10);

	$str="";
	for($j=0;$j<6;$j++){
		$subtemp=$temp & intval(base_convert("3d", 16, 10));
		$str.=$chars[$subtemp];
		$temp=$temp>>5;
	}
	unset($temp);
	$rs[]=$str;
}

print_r($rs);
