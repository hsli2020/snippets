<?php
/**
 * Writes a message to the log.
 * 
 * This method adds a message line to the log file defined by the config.
 * This includes the priority level, user IP, and a backtrace of the call.
 * @param string $msg <p>The message to write to the log.</p>
 * @param string $level [optional] <p>The priority level of the message.
 * This is merely for the benefit of the user and does not affect how
 * the code runs. The values used in this library are "Info", "Warning",
 * "Urgent", and "Throttle".</p>
 * @return boolean <b>FALSE</b> if the message is empty, NULL if logging is muted
 * @throws Exception If the file can't be written to.
 */
protected function log($msg, $level = 'Info'){
    if ($msg != false) {
        $backtrace = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS);
        
        if (file_exists($this->config)){
            include($this->config);
        } else {
            throw new Exception("Config file does not exist!");
        }
        if (isset($logfunction) && $logfunction != '' && function_exists($logfunction)){
            switch ($level){
               case('Info'): $loglevel = LOG_INFO; break; 
               case('Throttle'): $loglevel = LOG_INFO; break; 
               case('Warning'): $loglevel = LOG_NOTICE; break; 
               case('Urgent'): $loglevel = LOG_ERR; break; 
               default: $loglevel = LOG_INFO;
            }
            call_user_func($logfunction,$msg,$loglevel);
        }
        
        if (isset($muteLog) && $muteLog == true){
            return;
        }
        
        if(isset($userName) && $userName != ''){ 
            $name = $userName;
        }else{
            $name = 'guest';
        }
        
        if(isset($backtrace) && isset($backtrace[1]) && isset($backtrace[1]['file']) && isset($backtrace[1]['line']) && isset($backtrace[1]['function'])){
            $fileName = basename($backtrace[1]['file']);
            $file = $backtrace[1]['file'];
            $line = $backtrace[1]['line'];
            $function = $backtrace[1]['function'];
        }else{
            $fileName = basename($backtrace[0]['file']);
            $file = $backtrace[0]['file'];
            $line = $backtrace[0]['line'];
            $function = $backtrace[0]['function'];
        }
        if(isset($_SERVER['REMOTE_ADDR'])){
            $ip = $_SERVER['REMOTE_ADDR'];
            if($ip == '127.0.0.1')$ip = 'local';//save some char
        }else{
            $ip = 'cli';
        }
        if (!file_exists($this->logpath)) {
            //attempt to create the file if it does not exist
            file_put_contents($this->logpath, "This is the Amazon log, for Amazon classes to use.\n");
        }
        if (file_exists($this->logpath) && is_writable($this->logpath)){
            $str = "[$level][" . date("Y/m/d H:i:s") . " $name@$ip $fileName:$line $function] " . $msg;
            $fd = fopen($this->logpath, "a+");
            fwrite($fd,$str . "\r\n");
            fclose($fd);
        } else {
            throw new Exception('Error! Cannot write to log! ('.$this->logpath.')');
        }
    } else {
        return false;
    }
}
