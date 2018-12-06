<?php
/**
 * Batch Mockup
 *
 * I placed this in my BaseController.php file
 *
 * @usage
 *     $batch = new Batch('stats');
 *     $batch->columns = ['score', 'name'];
 *     $batch->data = [
 *         [1, 'john'],
 *         [4, 'fred'],
 *         [1, 'mickey'],
 *     ];
 *     $batch->insert();
 *
 */
class BatchInsert
{
	public $table = null;
	public $columns = array();
	public $data = array();

	public $insert_string = '';

	// --------------------------------------------------------------

	public function __construct($table = false)
	{
		if ($table) {
			$this->table = $table;
		}

		$di = Phalcon\DI::getDefault();
		$this->db = $di->get('db');
	}

	// --------------------------------------------------------------

	public function insert()
	{
		$this->_validate();
		$this->_stringify();

		//echo $this->insert_string;
		try {
			$result = $this->db->execute($this->insert_string);
		} catch (Exception $e) {
			echo $e->getMessage();
		}

		$this->insert_string = '';
		return $result;
	}

	// --------------------------------------------------------------

	private function _stringify()
	{
		$columns = sprintf('`%s`', implode('`,`', $this->columns));

		$str = '';
		foreach ($this->data as $values) {

			foreach ($values as &$val) {
				if (is_null($val)) {
					$val = 'NULL';
					continue;
				}
				if (is_string($val)) {
					$val = "'".$val."'";
				}
			}

			$str .= sprintf('(%s),', implode(',', $values));
		}
		$str = rtrim($str, ',');

		$query = sprintf("INSERT INTO `%s` (%s) VALUES %s",
			$this->table,
			$columns,
			$str
		);

		$this->insert_string = $query;
	}

	private function _makeValues($item)
	{
		print_r($item);
	}

	// --------------------------------------------------------------

	private function _validate()
	{
		if ($this->table == null) {
			throw new \Exception('Batch Table must be defined');
		}

		if (count($this->columns) == 0) {
			throw new \Exception('Batch Columns cannot be empty');
		}

		$required_count = count($this->columns);
		foreach ($this->data as $value) {
			if (count($value) !== $required_count) {
				throw new \Exception('Batch Data must match the same column count of ' . $required_count);
			}
		}
	}

	// --------------------------------------------------------------

}
