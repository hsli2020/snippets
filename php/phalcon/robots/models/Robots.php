<?php 

namespace Store\Toys;

use Phalcon\Mvc\Model;
use Phalcon\Validation;  # added this for new version of phalcon
use Phalcon\Mvc\Model\Message;
// use Phalcon\Mvc\Model\Validator\Uniqueness;  // working for old version of phalcon
// use Phalcon\Mvc\Model\Validator\InclusionIn; // working for old version of phalcon
use Phalcon\Validation\Validator\Uniqueness;
use Phalcon\Validation\Validator\InclusionIn;

class Robots extends Model
{
	public function validation()
	{

		// working for new version of Phalcon
		$validator = new Validation();

		$validator->add('name', new Uniqueness(
				[
					'message' => 'The robot name must be unique',
				]));

		$validator->add(
			'type',
			new InclusionIn(
				[
					'domain' => [
						'droid',
						'mechanical',
						'virtual',
					],
				]
			)
		);

		/** working for old version of phalcon

		// Type must be: droid, mechanical or virtual
		$this->validate(
			new InclusionIn(
				[
					'field' => 'type',
					'domain' => [
						'droid',
						'mechanical',
						'virtual',
					],
				]
			)
		);

		// Robots name must be unique
		$this->validate(
			new Uniqueness(
				[
					'field' => 'name',
					'message' => 'The robot name must be unique',
				]
			)
		);

		*/

		// Year cannot be less than zero
		if ($this->year < 0) {
			$this->appendMessage(
				new Message('The year cannot be less than zero')
			);
		};

		// Check if any messages have been produced
		if ($this->validationHasFailed() === true) {
			return false;
		}

		return $this->validate($validator);  // added this for new version of phalcon
	}
}