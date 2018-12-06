<?php

class AbTest
{
    protected $testName;
    protected $percentages;
    protected $minUserId;
    protected $maxUserId;

    public function __construct($testName, $percentages = array(50, 50))
    {
        $this->minUserId = null;
        $this->maxUserId = null;

        $this->testName = $testName;
        $this->percentages = $percentages;

        if (array_sum($percentages) != 100) {
            throw new InvalidArgumentException('Invalid groups setting');
        }
    }

    public function setMinUserId($userId)
    {
        $this->minUserId = $userId;
    }

    public function setMaxUserId($userId)
    {
        $this->maxUserId = $userId;
    }

    public function getGroup($userId)
    {
        if ((isset($this->minUserId) && ($userId < $this->minUserId)) ||
            (isset($this->maxUserId) && ($userId > $this->maxUserId))) {
            return 0;
        }

        $hash = substr($this->hash($userId), -3) * 0.1;
#       $hash = substr($this->hash($userId), -2) + 1;

        $sum = 0;
        foreach ($this->percentages as $group => $percent) {
            $sum += $percent;
            if ($hash <= $sum) {
                return $group + 1;
            }
        }

        return 0;  // should never reach here, but just in case
    }

    protected function hash($userId)
    {
        return sprintf('%u', crc32($userId . $this->testName));
    }
}
