<?php

class Config
{
  /**
   * Loads the section $section from the config file $filename.
   * If the $section is null, then all sections in the ini file are loaded.
   *
   * If the section name contains a ":" then the section name to the right
   * is loaded and included into the properties. Note that the keys in 
   * this $section will override any keys of the same name in the sections
   * that have been included via ":".
   * 
   * example ini file:
   *      [all]
   *      db.connection = database
   *      hostname = live
   *
   *      [staging : all]
   *      hostname = staging
   *
   * after calling $data = Config::load($file, 'staging'); then
   *      $data['hostname'] === "staging"
   *      $data['db.connection'] === "database"
   */
  public static function load($filename, $section = null)
  {
    $ini = parse_ini_file($filename, true);
    $preProcessed = array();

    foreach ($ini as $key => $data) {

      $bits = explode(':', $key);
      $numberOfBits = count($bits);
      $thisSection = trim($bits[0]);

      switch ($numberOfBits) {
      case 1:
        $preProcessed[$thisSection] = $data;
        break;

      case 2:
        $extendedSection = trim($bits[1]);
        $preProcessed[$thisSection] = array_merge(array(';extends'=>$extendedSection), $data);
        break;

      default:
        throw new Exception("Section '$thisSection' may not extend multiple sections in $filename");
      }
    }

    if (null === $section) {
      $array = array();
      foreach ($preProcessed as $sectionName => $sectionData) {
        $array[$sectionName] = self::_processExtends($preProcessed, $sectionName);
      }
    } elseif (is_array($section)) {
      $array = array();
      foreach ($section as $sectionName) {
        if (!isset($preProcessed[$sectionName])) {
          throw new Exception("Section '$sectionName' cannot be found in $filename");
        }
        $array = array_merge(self::_processExtends($preProcessed, $sectionName), $array);
      }
    } else {
      if (!isset($preProcessed[$section])) {
        throw new Exception("Section '$section' cannot be found in $filename");
      }
    }
    return $array;
  }

  protected static function _processExtends($ini, $section, $config = array())
  {
    $thisSection = $ini[$section];

    foreach ($thisSection as $key => $value) {
      if (strtolower($key) == ';extends') {
        if (isset($ini[$value])) {
          $config = self::_processExtends($ini, $value, $config);
        } else {
          throw new Exception("Section '$section' cannot be found");
        }
      } else {
        $config[$key] = $value;
      }
    }
    return $config;
  }
}

$ini = Config::load('a.ini');
print_r($ini);

