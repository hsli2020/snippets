<?php // User.php
/**
 * A Data Mapper, is a Data Access Layer that performs bidirectional transfer of data between
 * a persistent data store (often a relational database) and an in memory data representation
 * (the domain layer). The goal of the pattern is to keep the in memory representation and the
 * persistent data store independent of each other and the data mapper itself. The layer is
 * composed of one or more mappers (or Data Access Objects), performing the data transfer.
 * Mapper implementations vary in scope. Generic mappers will handle many different domain
 * entity types, dedicated mappers will handle one or a few.
 * 
 * The key point of this pattern is, unlike Active Record pattern, the data model follows
 * Single Responsibility Principle.
 * 
 * DB Object Relational Mapper (ORM) : Doctrine2 uses DAO named as “EntityRepository”
 */
namespace DesignPatterns\Structural\DataMapper;

/**
 * DataMapper pattern
 *
 * This is our representation of a DataBase record in the memory (Entity)
 *
 * Validation would also go in this object
 *
 */
class User
{
    /**
     * @var int
     */
    protected $userId;

    /**
     * @var string
     */
    protected $username;

    /**
     * @var string
     */
    protected $email;

    /**
     * @param null $id
     * @param null $username
     * @param null $email
     */
    public function __construct($id = null, $username = null, $email = null)
    {
        $this->userId = $id;
        $this->username = $username;
        $this->email = $email;
    }

    /**
     * @return int
     */
    public function getUserId()
    {
        return $this->userId;
    }

    /**
     * @param int $userId
     */
    public function setUserID($userId)
    {
        $this->userId = $userId;
    }

    /**
     * @return string
     */
    public function getUsername()
    {
        return $this->username;
    }

    /**
     * @param string $username
     */
    public function setUsername($username)
    {
        $this->username = $username;
    }

    /**
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * @param string $email
     */
    public function setEmail($email)
    {
        $this->email = $email;
    }
}


<?php  // UserMapper.php

namespace DesignPatterns\Structural\DataMapper;

/**
 * class UserMapper
 */
class UserMapper
{
    /**
     * @var DBAL
     */
    protected $adapter;

    /**
     * @param DBAL $dbLayer
     */
    public function __construct(DBAL $dbLayer)
    {
        $this->adapter = $dbLayer;
    }

    /**
     * saves a user object from memory to Database
     *
     * @param User $user
     *
     * @return boolean
     */
    public function save(User $user)
    {
        /* $data keys should correspond to valid Table columns on the Database */
        $data = array(
            'userid'   => $user->getUserId(),
            'username' => $user->getUsername(),
            'email'    => $user->getEmail(),
        );

        /* if no ID specified create new user else update the one in the Database */
        if (null === ($id = $user->getUserId())) {
            unset($data['userid']);
            $this->adapter->insert($data);

            return true;
        } else {
            $this->adapter->update($data, array('userid = ?' => $id));

            return true;
        }
    }

    /**
     * finds a user from Database based on ID and returns a User object located
     * in memory
     *
     * @param int $id
     *
     * @throws \InvalidArgumentException
     * @return User
     */
    public function findById($id)
    {
        $result = $this->adapter->find($id);

        if (0 == count($result)) {
            throw new \InvalidArgumentException("User #$id not found");
        }
        $row = $result->current();

        return $this->mapObject($row);
    }

    /**
     * fetches an array from Database and returns an array of User objects
     * located in memory
     *
     * @return array
     */
    public function findAll()
    {
        $resultSet = $this->adapter->findAll();
        $entries   = array();

        foreach ($resultSet as $row) {
            $entries[] = $this->mapObject($row);
        }

        return $entries;
    }

    /**
     * Maps a table row to an object
     *
     * @param array $row
     *
     * @return User
     */
    protected function mapObject(array $row)
    {
        $entry = new User();
        $entry->setUserID($row['userid']);
        $entry->setUsername($row['username']);
        $entry->setEmail($row['email']);

        return $entry;
    }
}
