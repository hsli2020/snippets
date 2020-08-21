<?php

# https://github.com/donatj/PhpUserAgent/blob/master/src/UserAgentParser.php

const PLATFORM        = 'platform';
const BROWSER         = 'browser';
const BROWSER_VERSION = 'version';

/**
 * Parses a user agent string into its important parts
 *
 * @param string|null $u_agent User agent string to parse or null. Uses $_SERVER['HTTP_USER_AGENT'] on NULL
 * @return string[] an array with 'browser', 'version' and 'platform' keys
 * @throws \InvalidArgumentException on not having a proper user agent to parse.
 */
function parse_user_agent($u_agent = null)
{
    if ($u_agent === null && isset($_SERVER['HTTP_USER_AGENT'])) {
        $u_agent = (string)$_SERVER['HTTP_USER_AGENT'];
    }

    if( $u_agent === null ) {
        throw new \InvalidArgumentException('parse_user_agent requires a user agent');
    }

    $platform = null;
    $browser  = null;
    $version  = null;

    $empty = array( PLATFORM => $platform, BROWSER => $browser, BROWSER_VERSION => $version );

    if( !$u_agent ) {
        return $empty;
    }

    if( preg_match('/\((.*?)\)/m', $u_agent, $parent_matches) ) {
        preg_match_all(<<<'REGEX'
/(?P<platform>BB\d+;|Android|CrOS|Tizen|iPhone|iPad|iPod|Linux|(Open|Net|Free)BSD|Macintosh|Windows(\ Phone)?|Silk|linux-gnu|BlackBerry|PlayBook|X11|(New\ )?Nintendo\ (WiiU?|3?DS|Switch)|Xbox(\ One)?)
(?:\ [^;]*)?
(?:;|$)/imx
REGEX
            , $parent_matches[1], $result);

        $priority = array( 'Xbox One', 'Xbox', 'Windows Phone', 'Tizen', 'Android', 'FreeBSD', 'NetBSD', 'OpenBSD', 'CrOS', 'X11' );

        $result[PLATFORM] = array_unique($result[PLATFORM]);
        if( count($result[PLATFORM]) > 1 ) {
            if( $keys = array_intersect($priority, $result[PLATFORM]) ) {
                $platform = reset($keys);
            } else {
                $platform = $result[PLATFORM][0];
            }
        } elseif( isset($result[PLATFORM][0]) ) {
            $platform = $result[PLATFORM][0];
        }
    }

    if( $platform == 'linux-gnu' || $platform == 'X11' ) {
        $platform = 'Linux';
    } elseif( $platform == 'CrOS' ) {
        $platform = 'Chrome OS';
    }

    preg_match_all(<<<'REGEX'
%(?P<browser>Camino|Kindle(\ Fire)?|Firefox|Iceweasel|IceCat|Safari|MSIE|Trident|AppleWebKit|
TizenBrowser|(?:Headless)?Chrome|YaBrowser|Vivaldi|IEMobile|Opera|OPR|Silk|Midori|Edge|Edg|CriOS|UCBrowser|Puffin|OculusBrowser|SamsungBrowser|
Baiduspider|Applebot|Googlebot|YandexBot|bingbot|Lynx|Version|Wget|curl|
Valve\ Steam\ Tenfoot|
NintendoBrowser|PLAYSTATION\ (\d|Vita)+)
(?:\)?;?)
(?:(?:[:/ ])(?P<version>[0-9A-Z.]+)|/(?:[A-Z]*))%ix
REGEX
        , $u_agent, $result);

    // If nothing matched, return null (to avoid undefined index errors)
    if( !isset($result[BROWSER][0]) || !isset($result[BROWSER_VERSION][0]) ) {
        if( preg_match('%^(?!Mozilla)(?P<browser>[A-Z0-9\-]+)(/(?P<version>[0-9A-Z.]+))?%ix', $u_agent, $result) ) {
            return array( PLATFORM => $platform ?: null, BROWSER => $result[BROWSER], BROWSER_VERSION => empty($result[BROWSER_VERSION]) ? null : $result[BROWSER_VERSION] );
        }

        return $empty;
    }

    if( preg_match('/rv:(?P<version>[0-9A-Z.]+)/i', $u_agent, $rv_result) ) {
        $rv_result = $rv_result[BROWSER_VERSION];
    }

    $browser = $result[BROWSER][0];
    $version = $result[BROWSER_VERSION][0];

    $lowerBrowser = array_map('strtolower', $result[BROWSER]);

    $find = function ( $search, &$key = null, &$value = null ) use ( $lowerBrowser ) {
        $search = (array)$search;

        foreach( $search as $val ) {
            $xkey = array_search(strtolower($val), $lowerBrowser);
            if( $xkey !== false ) {
                $value = $val;
                $key   = $xkey;

                return true;
            }
        }

        return false;
    };

    $findT = function ( array $search, &$key = null, &$value = null ) use ( $find ) {
        $value2 = null;
        if( $find(array_keys($search), $key, $value2) ) {
            $value = $search[$value2];

            return true;
        }

        return false;
    };

    $key = 0;
    $val = '';
    if( $findT(array( 'OPR' => 'Opera', 'UCBrowser' => 'UC Browser', 'YaBrowser' => 'Yandex', 'Iceweasel' => 'Firefox', 'Icecat' => 'Firefox', 'CriOS' => 'Chrome', 'Edg' => 'Edge' ), $key, $browser) ) {
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $find('Playstation Vita', $key, $platform) ) {
        $platform = 'PlayStation Vita';
        $browser  = 'Browser';
    } elseif( $find(array( 'Kindle Fire', 'Silk' ), $key, $val) ) {
        $browser  = $val == 'Silk' ? 'Silk' : 'Kindle';
        $platform = 'Kindle Fire';
        if( !($version = $result[BROWSER_VERSION][$key]) || !is_numeric($version[0]) ) {
            $version = $result[BROWSER_VERSION][array_search('Version', $result[BROWSER])];
        }
    } elseif( $find('NintendoBrowser', $key) || $platform == 'Nintendo 3DS' ) {
        $browser = 'NintendoBrowser';
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $find('Kindle', $key, $platform) ) {
        $browser = $result[BROWSER][$key];
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $find('Opera', $key, $browser) ) {
        $find('Version', $key);
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $find('Puffin', $key, $browser) ) {
        $version = $result[BROWSER_VERSION][$key];
        if( strlen($version) > 3 ) {
            $part = substr($version, -2);
            if( ctype_upper($part) ) {
                $version = substr($version, 0, -2);

                $flags = array( 'IP' => 'iPhone', 'IT' => 'iPad', 'AP' => 'Android', 'AT' => 'Android', 'WP' => 'Windows Phone', 'WT' => 'Windows' );
                if( isset($flags[$part]) ) {
                    $platform = $flags[$part];
                }
            }
        }
    } elseif( $find(array( 'Applebot', 'IEMobile', 'Edge', 'Midori', 'Vivaldi', 'OculusBrowser', 'SamsungBrowser', 'Valve Steam Tenfoot', 'Chrome', 'HeadlessChrome' ), $key, $browser) ) {
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $rv_result && $find('Trident') ) {
        $browser = 'MSIE';
        $version = $rv_result;
    } elseif( $browser == 'AppleWebKit' ) {
        if( $platform == 'Android' ) {
            $browser = 'Android Browser';
        } elseif( strpos($platform, 'BB') === 0 ) {
            $browser  = 'BlackBerry Browser';
            $platform = 'BlackBerry';
        } elseif( $platform == 'BlackBerry' || $platform == 'PlayBook' ) {
            $browser = 'BlackBerry Browser';
        } else {
            $find('Safari', $key, $browser) || $find('TizenBrowser', $key, $browser);
        }

        $find('Version', $key);
        $version = $result[BROWSER_VERSION][$key];
    } elseif( $pKey = preg_grep('/playstation \d/i', $result[BROWSER]) ) {
        $pKey = reset($pKey);

        $platform = 'PlayStation ' . preg_replace('/\D/', '', $pKey);
        $browser  = 'NetFront';
    }

    return array( PLATFORM => $platform ?: null, BROWSER => $browser ?: null, BROWSER_VERSION => $version ?: null );
}

$list = [
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.3;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.3;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89  Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83  Safari/537.36 Edg/85.0.564.41",
    "Mozilla/5.0 (Windows NT 6.3;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83  Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83  Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83  Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.3;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83  Safari/537.36 Edg/85.0.564.44",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.3;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.63",

    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; rv:11.0) like Gecko",

    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:79.0) Gecko/20100101 Firefox/79.0",
    "Mozilla/5.0 (Windows NT 6.2;  Win64; x64; rv:79.0) Gecko/20100101 Firefox/79.0",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64; rv:79.0) Gecko/20100101 Firefox/79.0",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
    "Mozilla/5.0 (Windows NT 6.2;  Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64; rv:75.0) Gecko/20100101 Firefox/75.0",
    "Mozilla/5.0 (Windows NT 6.1;  Win64; x64; rv:81.0) Gecko/20100101 Firefox/81.0",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:81.0) Gecko/20100101 Firefox/81.0",
    "Mozilla/5.0 (Windows NT 6.2;  Win64; x64; rv:81.0) Gecko/20100101 Firefox/81.0",

    "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; Info",

    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Safari/537.36",

    "Mozilla/5.0 (Linux; Android 10; SM-G965W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 10; SM-G965W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.81  Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 9;  CLT-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.81  Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 10; VOG-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.81  Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 10; SM-G965W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 9;  CLT-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 9;  SM-G950W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 10; VOG-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 9;  CLT-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.120 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 9;  CLT-L04)  AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",
    "Mozilla/5.0 (Linux; Android 10; SM-G965W) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.127 Mobile Safari/537.36",

    "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1 Mobile/15E148 Safari/604.1",

    "Mozilla/5.0 (iPhone; CPU iPhone OS 13_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.2 Mobile/15E148 Safari/604.1",
    "Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 Mobile/15E148 Safari/604.1",
];

foreach ($list as $ua) {
    $arr = parse_user_agent($ua);
    echo implode('/', $arr), PHP_EOL;
}
