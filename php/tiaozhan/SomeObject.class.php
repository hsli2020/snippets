<?php

/*
 * A class representing the object described in the README.
 */
class SomeObject {
	protected $statusLog;
	protected $startDate;
	protected $stopDate;

	public function __construct($statusLog, $startDate = null, $stopDate = null) {
		$this->statusLog = $statusLog;
		$this->startDate = $startDate;
		$this->stopDate = $stopDate;
	}

	/*
	 * @return array
	 */
	public function getStatusLog() {
		return $this->statusLog;
	}

	/*
	 * @return int|null
	 */
	public function getStartDate() {
		return $this->startDate;
	}

	/*
	 * @return int|null
	 */
	public function getStopDate() {
		return $this->stopDate;
	}

}

?>