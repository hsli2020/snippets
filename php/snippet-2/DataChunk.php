<?php

class DataChunk
{
    protected $db;
    protected $table;
    protected $columns;
    protected $data = [];
    protected $threshold;

    public function __construct($db, $table, $columns, $threshold = 1000)
    {
        $this->db        = $db;
        $this->table     = $table;
        $this->columns   = $columns;
        $this->threshold = $threshold;
    }

    public function setThreshold($threshold)
    {
        $this->threshold = $threshold;
    }

    public function add($row)
    {
        $this->data[] = $row;

        if (count($this->data) >= $this->threshold) {
            $this->save();
        }
    }

    public function save()
    {
        $sql = Toolkit\Utils::genInsertSql($this->table, $this->columns, $this->data);
        $this->db->execute($sql);
        $this->data = [];
    }
}
