<?php

// File: DomainObjectAbstract.php

abstract class DomainObjectAbstract
{
    protected $_data = array();

    public function __construct(array $data = NULL)
    {
        if ($data !== NULL) {
            // populate domain object with an array of data
            foreach ($data as $property => $value) {
                if (!empty($property)) {

                   // why the constructor operates on the object and not:
                   // $this->_data[$property] = $value; see 'this is why'

                   $this->$property = $value;
                }
            }
        }
    }

    // set domain object property
    public function __set($property, $value)
    {
        if (!array_key_exists($property, $this->_data)) {
            throw new ModelObjectException('The specified property is not valid for this domain object.');
        }

        if (strtolower($property) === 'id' AND $this->_data['id'] !== NULL) {
            throw new DomainObjectException('ID for this domain object is immutable.');
        }
        $this->_data[$property] = $value;
    }

    // get domain object property
    public function __get($property)
    {
        if (!array_key_exists($property, $this->_data)) {
            throw new DomainObjectException('The property requested is not valid for this domain object.');
        }
        return $this->_data[$property];
    }

    // check if given domain object property has been set
    public function __isset($property)
    {
        return isset($this->_data[$property]);
    }

    // unset domain object property
    public function __unset($property)
    {
        if (isset($this->_data[$property])) {
            unset($this->_data[$property]);
        }
    }
}

class DomainObjectException extends Exception {}

// File: User.php

class User extends DomainObjectAbstract
{
    // this is why
    protected $_data = array(
        'id'    => NULL,
        'fname' => '',
        'lname' => '',
        'email' => '',

        // property name, column name, value
        // 'FirstName' => array('fname', ''),
    );
}


$user = new User();
$user->fname = 'Susan';
$user->lname = 'Norton';
$user->email = 'susan@domain.com';


// File: MyApplication/Entity/AbstractEntity.php

namespace MyApplicationEntity;

abstract class AbstractEntity
{
    protected $_values = array();
    protected $_allowedFields = array();

    /**
     * Class constructor
     */
    public function __construct(array $data)
    {
        foreach ($data as $name => $value) {
            $this->$name = $value;
        }
    }

    /**
     * Assign a value to the specified field via the corresponding mutator (if it exists);
     * otherwise, assign the value directly to the '$_values' protected array
     */
    public function __set($name, $value)
    {
        if (!in_array($name, $this->_allowedFields)) {
            throw new EntityException('The field ' . $name . ' is not allowed for this entity.');
        }
        $mutator = 'set' . ucfirst($name);
        if (method_exists($this, $mutator) && is_callable(array($this, $mutator))) {
            $this->$mutator($value);
        }
        else {
            $this->_values[$name] = $value;
        }
    }

    /**
     * Get the value assigned to the specified field via the corresponding getter (if it exists);
     * otherwise, get the value directly from the '$_values' protected array
     */
    public function __get($name)
    {
        if (!in_array($name, $this->_allowedFields)) {
            throw new EntityException('The field ' . $name . ' is not allowed for this entity.');
        }
        $accessor = 'get' . ucfirst($name);
        if (method_exists($this, $accessor) && is_callable(array($this, $accessor))) {
            return $this->$accessor;
        }
        if (isset($this->_values[$name])) {
            return $this->_values[$name];
        }
        throw new EntityException('The field ' . $name . ' has not been set for this entity yet.');
    }

    /**
     * Check if the specified field has been assigned to the entity
     */
    public function __isset($name)
    {
        if (!in_array($name, $this->_allowedFields)) {
            throw new EntityException('The field ' . $name . ' is not allowed for this entity.');
        }
        return isset($this->_values[$name]);
    }

    /**
     * Unset the specified field from the entity
     */
    public function __unset($name)
    {
        if (!in_array($name, $this->_allowedFields)) {
            throw new EntityException('The field ' . $name . ' is not allowed for this entity.');
        }
        if (isset($this->_values[$name])) {
            unset($this->_values[$name]);
        }
    }

    /**
     * Get an associative array with the values assigned to the fields of the entity
     */
    public function toArray()
    {
        return $this->_values;
    }
}


// File: MyApplication/Entity/EntityException.php

namespace MyApplicationEntity;

class EntityException extends Exception { }


// File: MyApplication/Entity/User.php

namespace MyApplicationEntity;

class User extends AbstractEntity
{   
    protected $_allowedFields = array('id', 'fname', 'lname', 'email');
    
    /**
     * Set the user's ID
     */
    public function setId($id)
    {
        if(!filter_var($id, FILTER_VALIDATE_INT, array('options' => array('min_range' => 1, 'max_range' => 999999)))) {
            throw new EntityException('The specified ID is invalid.');
        }
        $this->_values['id'] = $id;
    }
    
    /**
     * Set the user's first name 
     */  
    public function setFname($fname)
    {
        if (strlen($fname) < 2 || strlen($fname) > 32) {
            throw new EntityException('The specified first name is invalid.');
        }
        $this->_values['fname'] = $fname;
    }
        
    /**
     * Set the user's last name
     */ 
    public function setLname($lname)
    {
        if (strlen($lname) < 2 || strlen($lname) > 32) {
            throw new EntityException('The specified last name is invalid.');
        }
        $this->_values['lname'] = $lname;
    }
    
    /**
     * Set the user's email address
     */ 
    public function setEmail($email)
    {
        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            throw new EntityException('The specified email address is invalid.');
        }
        $this->_values['email'] = $email;
    }                    
}


// File: MyApplication/Mapper/DataMapperInterface.php

namespace MyApplicationMapper;

interface DataMapperInterface
{
    public function findById($id);

    public function findAll();

    public function search($criteria);

    public function insert($entity);

    public function update($entity);

    public function delete($entity);
}


// File: MyApplication/Mapper/AbstractDataMapper.php

namespace MyApplicationMapper;

use MyApplicationDatabase, MyApplicationCollection;

abstract class AbstractDataMapper implements DataMapperInterface
{
    protected $_adapter;
    protected $_collection;
    protected $_entityTable;
    protected $_entityClass;

    /**
     * Constructor
     */
    public function  __construct(DatabaseDatabaseAdapterInterface $adapter, 
        CollectionAbstractCollection $collection, array $entityOptions = array())
    {
        $this->_adapter = $adapter;
        $this->_collection = $collection;
        if (isset($entityOptions['entityTable'])) {
            $this->setEntityTable($entityOptions['entityTable']);
        }
        if (isset($entityOptions['entityClass'])) {
            $this->setEntityClass($entityOptions['entityClass']);
        }
    }

    /**
     * Get the database adapter
     */
    public function getAdapter()
    {
        return $this->_adapter;
    }

    /**
     * Get the collection
     */
    public function getCollection()
    {
        return $this->_collection;
    }

    /**
     * Set the entity table
     */
    public function setEntityTable($entityTable)
    {
        if (!is_string($table) || empty ($entityTable)) {
            throw new DataMapperException('The specified entity table is invalid.');
        }
        $this->_entityTable = $entityTable;
    }

    /**
     * Get the entity table
     */
    public function getEntityTable()
    {
        return $this->_entityTable;
    }

    /**
     * Set the entity class
     */
    public function setEntityClass($entityClass)
    {
        if (!class_exists($entityClass)) {
            throw new DataMapperException('The specified entity class is invalid.');
        }
        $this->_entityClass = $entityClass;
    }

    /**
     * Get the entity class
     */
    public function getEntityClass()
    {
        return $this->_entityClass;
    }

    /**
     * Find an entity by its ID
     */
    public function findById($id)
    {
        $this->_adapter->select($this->_entityTable, "id = $id");
        if ($data = $this->_adapter->fetch()) {
            return new $this->_entityClass($data);
        }
        return null;
    }

    /**
     * Find all the entities
     */
    public function findAll()
    {
        $this->_adapter->select($this->_entityTable);
        while ($data = $this->_adapter->fetch($this->_entityTable)) {
            $this->_collection[] = new $this->_entityClass($data);
        }
        return $this->_collection;
    }

    /**
     * Find all the entities that match the specified criteria
     */
    public function search($criteria)
    {
        $this->_adapter->select($this->_entityTable, $criteria);
        while ($data = $this->_adapter->fetch()) {
            $this->_collection[] = new $this->_entityClass($data);
        }
        return $this->_collection;
    }
    
    /**
     * Insert a new row in the table corresponding to the specified entity
     */
    public function insert($entity)
    {
        if ($entity instanceof $this->_entityClass) {
            return $this->_adapter->insert($this->_entityTable, $entity->toArray());
        }
        throw new DataMapperException('The specified entity is not allowed for this mapper.');
    }

    /**
     * Update the row in the table corresponding to the specified entity
     */
    public function update($entity)
    {
        if ($entity instanceof $this->_entityClass) {
            $data = $entity->toArray();
            $id = $entity->id;
            unset($data['id']);
            return $this->_adapter->update($this->_entityTable, $data, "id = $id");
        }
        throw new DataMapperException('The specified entity is not allowed for this mapper.');
    }

    /**
     * Delete the row in the table corresponding to the specified entity or ID
     */
    public function delete($id)
    {
        if ($id instanceof $this->_entityClass) {
            $id = $id->id;
        }
        return $this->_adapter->delete($this->_entityTable, "id = $id");
    }
}

 
// File: MyApplication/Mapper/DataMapperException.php

namespace MyApplicationMapper;

class DataMapperException extends Exception {}


// File: MyApplication/Mapper/UserMapper.php

namespace MyApplicationMapper;
use MyApplicationDatabase, MyApplicationCollection;

class UserMapper extends AbstractDataMapper
{
    protected $_entityClass = 'MyApplicationEntityUser';
    protected $_entityTable = 'users';  
    
    /**
     * Constructor
     */
    public function __construct(DatabaseDatabaseAdapterInterface $adapter, 
        CollectionUserCollection $collection)
    {
        parent::__construct($adapter, $collection); 
    }   
}


// File: DomainObjectAbstract.php
 
abstract class DomainObjectAbstract
{
    protected $_data = array();
   
    public function __construct(array $data = NULL)
    {
        if ($data !== NULL) {
            // populate domain object with an array of data
            foreach ($data as $property => $value) {
                if (!empty($property)) {
                   $this->$property = $value;
                }
            }
        }
    }
   
    // set domain object property
    public function __set($property, $value)
    {
        if (!array_key_exists($property, $this->_data)) {
            throw new ModelObjectException('The specified property is not valid for this domain object.'); 
        }
        if (strtolower($property) === 'id' AND $this->_data['id'] !== NULL) {
            throw new DomainObjectException('ID for this domain object is immutable.');
        }
        $this->_data[$property] = $value;
    }
   
    // get domain object property
    public function __get($property)
    {
        if (!array_key_exists($property, $this->_data)) {
            throw new DomainObjectException('The property requested is not valid for this domain object.');
        }
        return $this->_data[$property];
    } 
   
    // check if given domain object property has been set
    public function __isset($property)
    {
        return isset($this->_data[$property]);
    }
   
    // unset domain object property
    public function __unset($property)
    {
        if (isset($this->_data[$property])) {
            unset($this->_data[$property]);
        }
    }
}
 
// File: DomainObjectException.php
 
class DomainObjectException extends Exception{}


// File: User.php

class User extends DomainObjectAbstract
{
    protected $_data = array(
        'id' => NULL, 
        'fname' => '', 
        'lname' => '', 
        'email' => ''
    );
}

$user = new User();
$user->fname = 'Julie';
$user->lname = 'Smith';
$user->email = 'julie@domain.com';


// File: MySQLAdapter.php

class MySQLAdapter
{
    private $_config = array();
    private static $_instance = NULL;
    private static $_connected = FALSE;
    private $_link = NULL;
    private $_result = NULL;
   
    // return Singleton instance of MySQLAdapter class
    public static function getInstance(array $config = array())
    {
        if (self::$_instance === NULL) {
            self::$_instance = new self($config);
        }
        return self::$_instance;
    }
   
    // private constructor
    private function __construct(array $config)
    {
        if (count($config) < 4) {
            throw new MySQLAdapterException('Invalid number of connection parameters');  
        }
        $this->_config = $config;
    }
   
    // prevent cloning class instance
    private function __clone() { }
   
    // connect to MySQL
    private function connect()
    {
        // connect only once
        if (self::$_connected === FALSE) {
            list($host, $user, $password, $database) = $this->_config;
            if ((!$this->_link = mysqli_connect($host, $user, $password, $database))) {
                throw new MySQLAdapterException('Error connecting to MySQL : ' 
                    . mysqli_connect_error());
            }
            self::$_connected = TRUE;
            unset($host, $user, $password, $database);      
        }
    }

    // perform query
    public function query($query)
    {
        if (is_string($query) and !empty($query)) {
            // lazy connect to MySQL
            $this->connect();
            if ((!$this->_result = mysqli_query($this->_link, $query))) {
                throw new MySQLAdapterException('Error performing query ' . 
                    $query . ' Error : ' . mysqli_error($this->_link));
            }
        }
    }
   
    // fetch row from result set
    public function fetch()
    {
        if ((!$row = mysqli_fetch_object($this->_result))) {
            mysqli_free_result($this->_result);
            return FALSE;
        }
        return $row;
    }

    // get insertion ID
    public function getInsertID()
    {
        if ($this->_link !== NUlL) {
            return mysqli_insert_id($this->_link); 
        }
        return NULL;  
    }
   
    // count rows in result set
    public function countRows()
    {
        if ($this->_result !== NULL) {
           return mysqli_num_rows($this->_result);
        }
        return 0;
    }
    
    // close the database connection
    function __destruct()
    {
        is_resource($this->_link) AND mysqli_close($this->_link);
    }
}
 
 
// File: MySQLAdapterException.php

class MySQLAdapterException extends Exception{}


// File: DataMapperAbstract.php
 
abstract class DataMapperAbstract
{
    protected $_db = NULL;
    protected $_table = '';
    protected $_identityMap = array();
   
    public function __construct(MySQLAdapter $db)
    {
        $this->_db = $db;   
    }
   
    // get domain object by ID (implemented by concrete domain object subclasses)
    abstract public function find($id);
   
    // insert/update domain object (implemented by concrete domain object subclasses)
    abstract public function save(DomainObjectAbstract $domainObject);
   
    // delete domain object (implemented by concrete domain object subclasses)
    abstract public function delete(DomainObjectAbstract $domainObject);
}


// File: UserMapper.php
 
class UserMapper extends DataMapperAbstract
{
    protected $_table = 'users';
   
    // fetch domain object by ID
    public function find($id)
    {
        // if the requested domain object exists in the identity map, get it from the there
        if (array_key_exists($id, $this->_identityMap)) {
            return $this->_identityMap[$id];
        }

        // if not, get domain object from the database
        $this->_db->query("SELECT * FROM $this->_table WHERE id = $id");
        if ($row = $this->_db->fetch()) {
            $user = new User;
            $user->id = $row->id;
            $user->fname = $row->fname;
            $user->lname = $row->lname;
            $user->email = $row->email;
            // save domain object to the identity map
            $this->_identityMap[$id] = $user;
            return $user;
        }
    }

    // save domain object
    public function save(DomainObjectAbstract $user)
    {
        // update domain object
        if ($user->id !== NULL) {
            $this->_db->query("UPDATE $this->_table SET 
                fname = '$user->fname', lname = '$user->lname', 
                email = '$user->email' WHERE id = $user->id");
        }
        // insert domain object
        else {
            $this->_db->query("INSERT INTO $this->_table (id, fname, lname, email) 
              VALUES (NULL, '$user->fname', '$user->lname', '$user->email')");
        }
    }
}


