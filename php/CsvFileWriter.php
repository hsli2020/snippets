<?php

const EOL = PHP_EOL;

class CsvFileWriter
{
    const MODE_CREATE = 0;
    const MODE_APPEND = 1;

    protected $filename;
    protected $filemode;
    protected $headline;
    protected $handle;

    public function __construct($filename = '', $headline = [])
    {
        $this->filename = $filename;
        $this->headline = $headline;
        $this->filemode = self::MODE_CREATE;
    }

    public function __destruct()
    {
        if (is_resource($this->handle)) {
            fclose($this->handle);
        }
    }

    public function setFilename($filename)
    {
        $this->filename = $filename;
        return $this;
    }

    public function setFilemode($filemode)
    {
        $this->filemode = $filemode;
        return $this;
    }

    public function setHeadline($headline)
    {
        $this->headline = $headline;
        return $this;
    }

    public function write($data, $filter = null)
    {
        if (!$this->handle) {
            if ($this->filemode == self::MODE_APPEND && file_exists($this->filename)) {
                $this->handle = fopen($this->filename, 'a');
            } else {
                $this->handle = fopen($this->filename, 'w');
                if ($this->headline) {
                    fputcsv($this->handle, $this->headline);
                }
            }
        }

        if (is_callable($filter)) {
            $data = $filter($data);
        }

        if (count($data) != count($this->headline)) {
            throw new \Exception('Wrong number of elements: '. var_export($data, true));
        }

        return fputcsv($this->handle, $data);
    }
}

#$columns = [ 'city', 'state', 'country' ];
#$csvfile = new CsvFileWriter('aaaaa.csv', $columns);
#//$csvfile->setFilemode(CsvFileWriter::MODE_APPEND);
#$csvfile->write([ 'Toronto',  'ON', 'CA' ]);
#$csvfile->write([ 'MONTREAL', 'QC', 'CA' ], function($d) {
#    $d[0] = strtolower($d[0]);
#    $d[1] = strtolower($d[1]);
#    $d[2] = strtolower($d[2]);
#    return $d;
#});
