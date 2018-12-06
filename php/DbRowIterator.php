<?php

# http://www.dragonbe.com/2015/07/speeding-up-database-calls-with-pdo-and.html

/*
$pdo = new \PDO(
    $config['db']['dsn'],
    $config['db']['username'],
    $config['db']['password']
);

$sql = 'SELECT * FROM `gen_contact` ORDER BY `contact_modified` DESC';

$stmt = $pdo->prepare($sql);
$stmt->execute();
$data = $stmt->fetchAll(\PDO::FETCH_OBJ);

echo 'Getting the contacts that changed the last 3 months' . PHP_EOL;
foreach ($data as $row) {
    $dt = new \DateTime('2015-04-01 00:00:00');
    if ($dt->format('Y-m-d') . '00:00:00' < $row->contact_modified) {
        echo sprintf(
            '%s (%s)| modified %s',
            $row->contact_name,
            $row->contact_email,
            $row->contact_modified
        ) . PHP_EOL;
    }
}
*/

/**
 * Class DbRowIterator
 *
 * File: Iterator/DbRowIterator.php
 */
class DbRowIterator implements Iterator
{
    /** @var \PDOStatement $pdoStatement The PDO Statement to execute */
    protected $pdoStatement;

    /** @var int $key The cursor pointer */
    protected $key;

    /** @var  bool|\stdClass The resultset for a single row */
    protected $result;

    /** @var  bool $valid Flag indicating there's a valid resource or not */
    protected $valid;

    public function __construct(\PDOStatement $PDOStatement)
    {
        $this->pdoStatement = $PDOStatement;
    }

    /**
     * @inheritDoc
     */
    public function current()
    {
        return $this->result;
    }

    /**
     * @inheritDoc
     */
    public function next()
    {
        $this->key++;
        $this->result = $this->pdoStatement->fetch(
            \PDO::FETCH_OBJ, 
            \PDO::FETCH_ORI_ABS, 
            $this->key
        );
        if (false === $this->result) {
            $this->valid = false;
            return null;
        }
    }

    /**
     * @inheritDoc
     */
    public function key()
    {
        return $this->key;
    }

    /**
     * @inheritDoc
     */
    public function valid()
    {
        return $this->valid;
    }

    /**
     * @inheritDoc
     */
    public function rewind()
    {
        $this->key = 0;
    }
}


class LastPeriodIterator extends FilterIterator
{
    protected $period;

    public function __construct(\Iterator $iterator, $period = 'last week')
    {
        parent::__construct($iterator);
        $this->period = $period;
    }
    public function accept()
    {
        if (!$this->getInnerIterator()->valid()) {
            return false;
        }
        $row = $this->getInnerIterator()->current();
        $dt = new \DateTime($this->period);
        if ($dt->format('Y-m-d') . '00:00:00' < $row->contact_modified) {
            return true;
        }
        return false;
    }
}


$pdo = new \PDO(
    $config['db']['dsn'],
    $config['db']['username'],
    $config['db']['password']
);

$sql = 'SELECT * FROM `gen_contact` ORDER BY `contact_modified` DESC';
$stmt = $pdo->prepare($sql, [\PDO::ATTR_CURSOR => \PDO::CURSOR_SCROLL]);
$stmt->execute();

$data = new DbRowIterator($stmt);
echo 'Getting the contacts that changed the last 3 months' . PHP_EOL;
$lastPeriod = new LastPeriodIterator($data, '2015-04-01 00:00:00');
foreach ($lastPeriod as $row) {
    echo sprintf(
        '%s (%s)| modified %s',
        $row->contact_name,
        $row->contact_email,
        $row->contact_modified
    ) . PHP_EOL;
}

