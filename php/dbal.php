<?php

/**
 * Connection interface.
 * Driver connections must implement this interface.
 *
 * This resembles (a subset of) the PDO interface.
 *
 * @since 2.0
 */
interface Doctrine\DBAL\Driver\Connection
{
    function prepare($prepareString);
    function query();
    function quote($input, $type=\PDO::PARAM_STR);
    function exec($statement);
    function lastInsertId($name = null);
    function beginTransaction();
    function commit();
    function rollBack();
    function errorCode();
    function errorInfo();
}

/**
 * A wrapper around a Doctrine\DBAL\Driver\Connection that adds features like
 * events, transaction isolation levels, configuration, emulated transaction nesting,
 * lazy connecting and more.
 */
class Doctrine\DBAL\Connection implements Doctrine\DBAL\Driver\Connection
{
    protected $_conn;            /** @var Doctrine\DBAL\Driver\Connection */
    protected $_config;          /** @var Doctrine\DBAL\Configuration */
    protected $_eventManager;    /** @var Doctrine\Common\EventManager */
    protected $_expr;            /** @var Doctrine\DBAL\Query\ExpressionBuilder */

    private $_isConnected = false;  /** Whether or not a connection has been established. */
    private $_transactionNestingLevel = 0;  /** The transaction nesting level. */
    private $_transactionIsolationLevel;  /** The currently active transaction isolation level. */
    private $_nestTransactionsWithSavepoints;  /** If nested transations should use savepoints */
    private $_params = array();  /** The parameters used during creation of the Connection instance. */

    /**
     * The DatabasePlatform object that provides information about the
     * database platform used by the connection.
     *
     * @var Doctrine\DBAL\Platforms\AbstractPlatform
     */
    protected $_platform;
    protected $_schemaManager;  /** @var Doctrine\DBAL\Schema\SchemaManager */
    protected $_driver;  /** @var Doctrine\DBAL\Driver */

    /** Flag that indicates whether the current transaction is marked for rollback only. */
    private $_isRollbackOnly = false;

    public function getParams()
    public function getDatabase()
    public function getHost()
    public function getPort()
    public function getUsername()
    public function getPassword()
    public function getDriver()
    public function getConfiguration()
    public function getEventManager()
    public function getDatabasePlatform()
    public function getExpressionBuilder()

    public function connect()

    public function fetchAssoc($statement, array $params = array())
    {
        return $this->executeQuery($statement, $params)->fetch(PDO::FETCH_ASSOC);
    }

    public function fetchAll($sql, array $params = array())
    public function fetchArray($statement, array $params = array())
    public function fetchColumn($statement, array $params = array(), $colnum = 0)

    public function isConnected()
    public function isTransactionActive()

    public function delete($tableName, array $identifier)

    public function close()

    public function setTransactionIsolation($level)
    public function getTransactionIsolation()

    public function update($tableName, array $data, array $identifier, array $types = array())
    public function insert($tableName, array $data, array $types = array())

    public function setCharset($charset)

    public function quoteIdentifier($str)
    public function quote($input, $type = null)

    /**
     * Prepares an SQL statement.
     *
     * @param string $statement The SQL statement to prepare.
     * @return Doctrine\DBAL\Driver\Statement The prepared statement.
     */
    public function prepare($statement)
    {
        $this->connect();
        return new Statement($statement, $this);
    }

    /**
     * Executes an, optionally parameterized, SQL query.
     *
     * If the query is parameterized, a prepared statement is used.
     * If an SQLLogger is configured, the execution is logged.
     *
     * @param string $query The SQL query to execute.
     * @param array $params The parameters to bind to the query, if any.
     * @param array $types The types the previous parameters are in.
     * @param QueryCacheProfile $qcp
     * @return Doctrine\DBAL\Driver\Statement The executed statement.
     * @internal PERF: Directly prepares a driver statement, not a wrapper.
     */
    public function executeQuery($query, array $params = array(), $types = array(), QueryCacheProfile $qcp = null)
    {
        if ($qcp !== null) {
            return $this->executeCacheQuery($query, $params, $types, $qcp);
        }

        $this->connect();

        $hasLogger = $this->_config->getSQLLogger() !== null;
        if ($hasLogger) {
            $this->_config->getSQLLogger()->startQuery($query, $params, $types);
        }

        if ($params) {
            list($query, $params, $types) = SQLParserUtils::expandListParameters($query, $params, $types);

            $stmt = $this->_conn->prepare($query);
            if ($types) {
                $this->_bindTypedValues($stmt, $params, $types);
                $stmt->execute();
            } else {
                $stmt->execute($params);
            }
        } else {
            $stmt = $this->_conn->query($query);
        }

        if ($hasLogger) {
            $this->_config->getSQLLogger()->stopQuery();
        }

        return $stmt;
    }

    /**
     * Execute a caching query and
     *
     * @param string $query
     * @param array $params
     * @param array $types
     * @param QueryCacheProfile $qcp
     * @return \Doctrine\DBAL\Driver\ResultStatement
     */
    public function executeCacheQuery($query, $params, $types, QueryCacheProfile $qcp)
    {
        $resultCache = $qcp->getResultCacheDriver() ?: $this->_config->getResultCacheImpl();
        if (!$resultCache) {
            throw CacheException::noResultDriverConfigured();
        }

        list($cacheKey, $realKey) = $qcp->generateCacheKeys($query, $params, $types);

        // fetch the row pointers entry
        if ($data = $resultCache->fetch($cacheKey)) {
            // is the real key part of this row pointers map or is the cache only pointing to other cache keys?
            if (isset($data[$realKey])) {
                return new ArrayStatement($data[$realKey]);
            } else if (array_key_exists($realKey, $data)) {
                return new ArrayStatement(array());
            }
        }
        return new ResultCacheStatement($this->executeQuery($query, $params, $types), $resultCache, $cacheKey, $realKey, $qcp->getLifetime());
    }

    /**
     * Executes an, optionally parameterized, SQL query and returns the result,
     * applying a given projection/transformation function on each row of the result.
     *
     * @param string $query The SQL query to execute.
     * @param array $params The parameters, if any.
     * @param Closure $mapper The transformation function that is applied on each row.
     *                        The function receives a single paramater, an array, that
     *                        represents a row of the result set.
     * @return mixed The projected result of the query.
     */
    public function project($query, array $params, Closure $function)
    {
        $result = array();
        $stmt = $this->executeQuery($query, $params ?: array());

        while ($row = $stmt->fetch()) {
            $result[] = $function($row);
        }

        $stmt->closeCursor();

        return $result;
    }

    public function query()
    public function executeUpdate($query, array $params = array(), array $types = array())
    public function exec($statement)

    public function getTransactionNestingLevel()

    public function errorCode()
    public function errorInfo()

    public function lastInsertId($seqName = null)

    public function setNestTransactionsWithSavepoints($nestTransactionsWithSavepoints)
    public function getNestTransactionsWithSavepoints()

    protected function _getNestedTransactionSavePointName()

    public function transactional(Closure $func)
    public function beginTransaction()
    public function commit()
    public function rollback()

    public function createSavepoint($savepoint)
    public function releaseSavepoint($savepoint)
    public function rollbackSavepoint($savepoint)

    public function getWrappedConnection()

    /**
     * Gets the SchemaManager that can be used to inspect or change the
     * database schema through the connection.
     *
     * @return Doctrine\DBAL\Schema\AbstractSchemaManager
     */
    public function getSchemaManager()
    {
        if ( ! $this->_schemaManager) {
            $this->_schemaManager = $this->_driver->getSchemaManager($this);
        }

        return $this->_schemaManager;
    }

    public function setRollbackOnly()
    public function isRollbackOnly()

    /**
     * Converts a given value to its database representation according to the conversion
     * rules of a specific DBAL mapping type.
     *
     * @param mixed $value The value to convert.
     * @param string $type The name of the DBAL mapping type.
     * @return mixed The converted value.
     */
    public function convertToDatabaseValue($value, $type)
    {
        return Type::getType($type)->convertToDatabaseValue($value, $this->_platform);
    }

    /**
     * Converts a given value to its PHP representation according to the conversion
     * rules of a specific DBAL mapping type.
     *
     * @param mixed $value The value to convert.
     * @param string $type The name of the DBAL mapping type.
     * @return mixed The converted type.
     */
    public function convertToPHPValue($value, $type)
    {
        return Type::getType($type)->convertToPHPValue($value, $this->_platform);
    }

    /**
     * Binds a set of parameters, some or all of which are typed with a PDO binding type
     * or DBAL mapping type, to a given statement.
     *
     * @param $stmt The statement to bind the values to.
     * @param array $params The map/list of named/positional parameters.
     * @param array $types The parameter types (PDO binding types or DBAL mapping types).
     * @internal Duck-typing used on the $stmt parameter to support driver statements as well as
     *           raw PDOStatement instances.
     */
    private function _bindTypedValues($stmt, array $params, array $types)
    {
        // Check whether parameters are positional or named. Mixing is not allowed, just like in PDO.
        if (is_int(key($params))) {
            // Positional parameters
            $typeOffset = array_key_exists(0, $types) ? -1 : 0;
            $bindIndex = 1;
            foreach ($params as $position => $value) {
                $typeIndex = $bindIndex + $typeOffset;
                if (isset($types[$typeIndex])) {
                    $type = $types[$typeIndex];
                    list($value, $bindingType) = $this->getBindingInfo($value, $type);
                    $stmt->bindValue($bindIndex, $value, $bindingType);
                } else {
                    $stmt->bindValue($bindIndex, $value);
                }
                ++$bindIndex;
            }
        } else {
            // Named parameters
            foreach ($params as $name => $value) {
                if (isset($types[$name])) {
                    $type = $types[$name];
                    list($value, $bindingType) = $this->getBindingInfo($value, $type);
                    $stmt->bindValue($name, $value, $bindingType);
                } else {
                    $stmt->bindValue($name, $value);
                }
            }
        }
    }

    /**
     * Gets the binding type of a given type. The given type can be a PDO or DBAL mapping type.
     *
     * @param mixed $value The value to bind
     * @param mixed $type The type to bind (PDO or DBAL)
     * @return array [0] => the (escaped) value, [1] => the binding type
     */
    private function getBindingInfo($value, $type)
    {
        if (is_string($type)) {
            $type = Type::getType($type);
        }
        if ($type instanceof Type) {
            $value = $type->convertToDatabaseValue($value, $this->_platform);
            $bindingType = $type->getBindingType();
        } else {
            $bindingType = $type; // PDO::PARAM_* constants
        }
        return array($value, $bindingType);
    }

    public function createQueryBuilder()
}

/**
 * QueryBuilder class is responsible to dynamically create SQL queries.
 *
 * Important: Verify that every feature you use will work with your database vendor.
 * SQL Query Builder does not attempt to validate the generated SQL at all.
 *
 * The query builder does no validation whatsoever if certain features even work with the
 * underlying database vendor. Limit queries and joins are NOT applied to UPDATE and DELETE statements
 * even if some vendors such as MySQL support it.
 */
class Doctrine\DBAL\Query\QueryBuilder
{
    /* The query types. */
    const SELECT = 0;
    const DELETE = 1;
    const UPDATE = 2;

    /** The builder states. */
    const STATE_DIRTY = 0;
    const STATE_CLEAN = 1;

    /**
     * @var Doctrine\DBAL\Connection DBAL Connection
     */
    private $connection = null;

    /**
     * @var array The array of SQL parts collected.
     */
    private $sqlParts = array(
        'select'  => array(),
        'from'    => array(),
        'join'    => array(),
        'set'     => array(),
        'where'   => null,
        'groupBy' => array(),
        'having'  => null,
        'orderBy' => array()
    );

    /**
     * @var string The complete SQL string for this query.
     */
    private $sql;

    /**
     * @var array The query parameters.
     */
    private $params = array();

    /**
     * @var array The parameter type map of this query.
     */
    private $paramTypes = array();

    /**
     * @var integer The type of query this is. Can be select, update or delete.
     */
    private $type = self::SELECT;

    /**
     * @var integer The state of the query object. Can be dirty or clean.
     */
    private $state = self::STATE_CLEAN;

    /**
     * @var integer The index of the first result to retrieve.
     */
    private $firstResult = null;

    /**
     * @var integer The maximum number of results to retrieve.
     */
    private $maxResults = null;

    /**
     * The counter of bound parameters used with (@see bindValue)
     *
     * @var int
     */
    private $boundCounter = 0;

    /**
     * Initializes a new <tt>QueryBuilder</tt>.
     *
     * @param Doctrine\DBAL\Connection $connection DBAL Connection
     */
    public function __construct(Connection $connection)
    {
        $this->connection = $connection;
    }

    /**
     * Gets an ExpressionBuilder used for object-oriented construction of query expressions.
     * This producer method is intended for convenient inline usage. Example:
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u')
     *         ->from('users', 'u')
     *         ->where($qb->expr()->eq('u.id', 1));
     * </code>
     *
     * For more complex expression construction, consider storing the expression
     * builder object in a local variable.
     *
     * @return Doctrine\DBAL\Query\ExpressionBuilder
     */
    public function expr() { return $this->connection->getExpressionBuilder(); }

    /**
     * Get the type of the currently built query.
     * @return integer
     */
    public function getType() { return $this->type; }

    /**
     * Get the associated DBAL Connection for this query builder.
     * @return Doctrine\DBAL\Connection
     */
    public function getConnection() { return $this->connection; }

    /**
     * Get the state of this query builder instance.
     * @return integer Either QueryBuilder::STATE_DIRTY or QueryBuilder::STATE_CLEAN.
     */
    public function getState() { return $this->state; }

    /**
     * Execute this query using the bound parameters and their types.
     *
     * Uses {@see Connection::executeQuery} for select statements and {@see Connection::executeUpdate}
     * for insert, update and delete statements.
     *
     * @return mixed
     */
    public function execute()
    {
        if ($this->type == self::SELECT) {
            return $this->connection->executeQuery($this->getSQL(), $this->params, $this->paramTypes);
        } else {
            return $this->connection->executeUpdate($this->getSQL(), $this->params, $this->paramTypes);
        }
    }

    /**
     * Get the complete SQL string formed by the current specifications of this QueryBuilder.
     *
     * <code>
     *     $qb = $em->createQueryBuilder()
     *         ->select('u')
     *         ->from('User', 'u')
     *     echo $qb->getSQL(); // SELECT u FROM User u
     * </code>
     *
     * @return string The sql query string.
     */
    public function getSQL()
    {
        if ($this->sql !== null && $this->state === self::STATE_CLEAN) {
            return $this->sql;
        }

        $sql = '';

        switch ($this->type) {
            case self::DELETE:
                $sql = $this->getSQLForDelete();
                break;

            case self::UPDATE:
                $sql = $this->getSQLForUpdate();
                break;

            case self::SELECT:
            default:
                $sql = $this->getSQLForSelect();
                break;
        }

        $this->state = self::STATE_CLEAN;
        $this->sql = $sql;

        return $sql;
    }

    /**
     * Sets a query parameter for the query being constructed.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u')
     *         ->from('users', 'u')
     *         ->where('u.id = :user_id')
     *         ->setParameter(':user_id', 1);
     * </code>
     *
     * @param string|integer $key The parameter position or name.
     * @param mixed $value The parameter value.
     * @param string|null $type PDO::PARAM_*
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function setParameter($key, $value, $type = null)
    {
        if ($type !== null) {
            $this->paramTypes[$key] = $type;
        }

        $this->params[$key] = $value;

        return $this;
    }

    /**
     * Sets a collection of query parameters for the query being constructed.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u')
     *         ->from('users', 'u')
     *         ->where('u.id = :user_id1 OR u.id = :user_id2')
     *         ->setParameters(array(
     *             ':user_id1' => 1,
     *             ':user_id2' => 2
     *         ));
     * </code>
     *
     * @param array $params The query parameters to set.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function setParameters(array $params, array $types = array())
    {
        $this->paramTypes = $types;
        $this->params = $params;

        return $this;
    }

    /**
     * Gets all defined query parameters for the query being constructed.
     * @return array The currently defined query parameters.
     */
    public function getParameters() { return $this->params; }

    /**
     * Gets a (previously set) query parameter of the query being constructed.
     *
     * @param mixed $key The key (index or name) of the bound parameter.
     * @return mixed The value of the bound parameter.
     */
    public function getParameter($key)
    {
        return isset($this->params[$key]) ? $this->params[$key] : null;
    }

    /**
     * Sets the position of the first result to retrieve (the "offset").
     *
     * @param integer $firstResult The first result to return.
     * @return Doctrine\DBAL\Query\QueryBuilder This QueryBuilder instance.
     */
    public function setFirstResult($firstResult)
    {
        $this->state = self::STATE_DIRTY;
        $this->firstResult = $firstResult;
        return $this;
    }

    /**
     * Gets the position of the first result the query object was set to retrieve (the "offset").
     * Returns NULL if {@link setFirstResult} was not applied to this QueryBuilder.
     *
     * @return integer The position of the first result.
     */
    public function getFirstResult() { return $this->firstResult; }

    /**
     * Sets the maximum number of results to retrieve (the "limit").
     *
     * @param integer $maxResults The maximum number of results to retrieve.
     * @return Doctrine\DBAL\Query\QueryBuilder This QueryBuilder instance.
     */
    public function setMaxResults($maxResults)
    {
        $this->state = self::STATE_DIRTY;
        $this->maxResults = $maxResults;
        return $this;
    }

    /**
     * Gets the maximum number of results the query object was set to retrieve (the "limit").
     * Returns NULL if {@link setMaxResults} was not applied to this query builder.
     *
     * @return integer Maximum number of results.
     */
    public function getMaxResults() { return $this->maxResults; }

    /**
     * Either appends to or replaces a single, generic query part.
     *
     * The available parts are: 'select', 'from', 'set', 'where',
     * 'groupBy', 'having' and 'orderBy'.
     *
     * @param string $sqlPartName
     * @param string $sqlPart
     * @param string $append
     * @return Doctrine\DBAL\Query\QueryBuilder This QueryBuilder instance.
     */
    public function add($sqlPartName, $sqlPart, $append = false)
    {
        $isArray = is_array($sqlPart);
        $isMultiple = is_array($this->sqlParts[$sqlPartName]);

        if ($isMultiple && !$isArray) {
            $sqlPart = array($sqlPart);
        }

        $this->state = self::STATE_DIRTY;

        if ($append) {
            if ($sqlPartName == "orderBy" || $sqlPartName == "groupBy" || $sqlPartName == "select" || $sqlPartName == "set") {
                foreach ($sqlPart AS $part) {
                    $this->sqlParts[$sqlPartName][] = $part;
                }
            } else if ($isArray && is_array($sqlPart[key($sqlPart)])) {
                $key = key($sqlPart);
                $this->sqlParts[$sqlPartName][$key][] = $sqlPart[$key];
            } else if ($isMultiple) {
                $this->sqlParts[$sqlPartName][] = $sqlPart;
            } else {
                $this->sqlParts[$sqlPartName] = $sqlPart;
            }

            return $this;
        }

        $this->sqlParts[$sqlPartName] = $sqlPart;

        return $this;
    }

    /**
     * Specifies an item that is to be returned in the query result.
     * Replaces any previously specified selections, if any.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.id', 'p.id')
     *         ->from('users', 'u')
     *         ->leftJoin('u', 'phonenumbers', 'p', 'u.id = p.user_id');
     * </code>
     *
     * @param mixed $select The selection expressions.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function select($select = null)
    {
        $this->type = self::SELECT;

        if (empty($select)) {
            return $this;
        }

        $selects = is_array($select) ? $select : func_get_args();

        return $this->add('select', $selects, false);
    }

    /**
     * Adds an item that is to be returned in the query result.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.id')
     *         ->addSelect('p.id')
     *         ->from('users', 'u')
     *         ->leftJoin('u', 'phonenumbers', 'u.id = p.user_id');
     * </code>
     *
     * @param mixed $select The selection expression.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function addSelect($select = null)
    {
        $this->type = self::SELECT;

        if (empty($select)) {
            return $this;
        }

        $selects = is_array($select) ? $select : func_get_args();

        return $this->add('select', $selects, true);
    }

    /**
     * Turns the query being built into a bulk delete query that ranges over
     * a certain table.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->delete('users', 'u')
     *         ->where('u.id = :user_id');
     *         ->setParameter(':user_id', 1);
     * </code>
     *
     * @param string $delete The table whose rows are subject to the deletion.
     * @param string $alias The table alias used in the constructed query.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function delete($delete = null, $alias = null)
    {
        $this->type = self::DELETE;

        if ( ! $delete) {
            return $this;
        }

        return $this->add('from', array(
            'table' => $delete,
            'alias' => $alias
        ));
    }

    /**
     * Turns the query being built into a bulk update query that ranges over
     * a certain table
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->update('users', 'u')
     *         ->set('u.password', md5('password'))
     *         ->where('u.id = ?');
     * </code>
     *
     * @param string $update The table whose rows are subject to the update.
     * @param string $alias The table alias used in the constructed query.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function update($update = null, $alias = null)
    {
        $this->type = self::UPDATE;

        if ( ! $update) {
            return $this;
        }

        return $this->add('from', array(
            'table' => $update,
            'alias' => $alias
        ));
    }

    /**
     * Create and add a query root corresponding to the table identified by the
     * given alias, forming a cartesian product with any existing query roots.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.id')
     *         ->from('users', 'u')
     * </code>
     *
     * @param string $from   The table
     * @param string $alias  The alias of the table
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function from($from, $alias)
    {
        return $this->add('from', array(
            'table' => $from,
            'alias' => $alias
        ), true);
    }

    /**
     * Creates and adds a join to the query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->join('u', 'phonenumbers', 'p', 'p.is_primary = 1');
     * </code>
     *
     * @param string $fromAlias The alias that points to a from clause
     * @param string $join The table name to join
     * @param string $alias The alias of the join table
     * @param string $condition The condition for the join
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function join($fromAlias, $join, $alias, $condition = null)
    {
        return $this->innerJoin($fromAlias, $join, $alias, $condition);
    }

    /**
     * Creates and adds a join to the query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->innerJoin('u', 'phonenumbers', 'p', 'p.is_primary = 1');
     * </code>
     *
     * @param string $fromAlias The alias that points to a from clause
     * @param string $join The table name to join
     * @param string $alias The alias of the join table
     * @param string $condition The condition for the join
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function innerJoin($fromAlias, $join, $alias, $condition = null)
    {
        return $this->add('join', array(
            $fromAlias => array(
                'joinType'      => 'inner',
                'joinTable'     => $join,
                'joinAlias'     => $alias,
                'joinCondition' => $condition
            )
        ), true);
    }

    /**
     * Creates and adds a left join to the query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->leftJoin('u', 'phonenumbers', 'p', 'p.is_primary = 1');
     * </code>
     *
     * @param string $fromAlias The alias that points to a from clause
     * @param string $join The table name to join
     * @param string $alias The alias of the join table
     * @param string $condition The condition for the join
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function leftJoin($fromAlias, $join, $alias, $condition = null)
    {
        return $this->add('join', array(
            $fromAlias => array(
                'joinType'      => 'left',
                'joinTable'     => $join,
                'joinAlias'     => $alias,
                'joinCondition' => $condition
            )
        ), true);
    }

    /**
     * Creates and adds a right join to the query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->rightJoin('u', 'phonenumbers', 'p', 'p.is_primary = 1');
     * </code>
     *
     * @param string $fromAlias The alias that points to a from clause
     * @param string $join The table name to join
     * @param string $alias The alias of the join table
     * @param string $condition The condition for the join
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function rightJoin($fromAlias, $join, $alias, $condition = null)
    {
        return $this->add('join', array(
            $fromAlias => array(
                'joinType'      => 'right',
                'joinTable'     => $join,
                'joinAlias'     => $alias,
                'joinCondition' => $condition
            )
        ), true);
    }

    /**
     * Sets a new value for a column in a bulk update query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->update('users', 'u')
     *         ->set('u.password', md5('password'))
     *         ->where('u.id = ?');
     * </code>
     *
     * @param string $key The column to set.
     * @param string $value The value, expression, placeholder, etc.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function set($key, $value)
    {
        return $this->add('set', $key .' = ' . $value, true);
    }

    /**
     * Specifies one or more restrictions to the query result.
     * Replaces any previously specified restrictions, if any.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->where('u.id = ?');
     *
     *     // You can optionally programatically build and/or expressions
     *     $qb = $conn->createQueryBuilder();
     *
     *     $or = $qb->expr()->orx();
     *     $or->add($qb->expr()->eq('u.id', 1));
     *     $or->add($qb->expr()->eq('u.id', 2));
     *
     *     $qb->update('users', 'u')
     *         ->set('u.password', md5('password'))
     *         ->where($or);
     * </code>
     *
     * @param mixed $predicates The restriction predicates.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function where($predicates)
    {
        if ( ! (func_num_args() == 1 && $predicates instanceof CompositeExpression) ) {
            $predicates = new CompositeExpression(CompositeExpression::TYPE_AND, func_get_args());
        }

        return $this->add('where', $predicates);
    }

    /**
     * Adds one or more restrictions to the query results, forming a logical
     * conjunction with any previously specified restrictions.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u')
     *         ->from('users', 'u')
     *         ->where('u.username LIKE ?')
     *         ->andWhere('u.is_active = 1');
     * </code>
     *
     * @param mixed $where The query restrictions.
     * @return QueryBuilder This QueryBuilder instance.
     * @see where()
     */
    public function andWhere($where)
    {
        $where = $this->getQueryPart('where');
        $args = func_get_args();

        if ($where instanceof CompositeExpression && $where->getType() === CompositeExpression::TYPE_AND) {
            $where->addMultiple($args);
        } else {
            array_unshift($args, $where);
            $where = new CompositeExpression(CompositeExpression::TYPE_AND, $args);
        }

        return $this->add('where', $where, true);
    }

    /**
     * Adds one or more restrictions to the query results, forming a logical
     * disjunction with any previously specified restrictions.
     *
     * <code>
     *     $qb = $em->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->where('u.id = 1')
     *         ->orWhere('u.id = 2');
     * </code>
     *
     * @param mixed $where The WHERE statement
     * @return QueryBuilder $qb
     * @see where()
     */
    public function orWhere($where)
    {
        $where = $this->getQueryPart('where');
        $args = func_get_args();

        if ($where instanceof CompositeExpression && $where->getType() === CompositeExpression::TYPE_OR) {
            $where->addMultiple($args);
        } else {
            array_unshift($args, $where);
            $where = new CompositeExpression(CompositeExpression::TYPE_OR, $args);
        }

        return $this->add('where', $where, true);
    }

    /**
     * Specifies a grouping over the results of the query.
     * Replaces any previously specified groupings, if any.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->groupBy('u.id');
     * </code>
     *
     * @param mixed $groupBy The grouping expression.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function groupBy($groupBy)
    {
        if (empty($groupBy)) {
            return $this;
        }

        $groupBy = is_array($groupBy) ? $groupBy : func_get_args();

        return $this->add('groupBy', $groupBy, false);
    }


    /**
     * Adds a grouping expression to the query.
     *
     * <code>
     *     $qb = $conn->createQueryBuilder()
     *         ->select('u.name')
     *         ->from('users', 'u')
     *         ->groupBy('u.lastLogin');
     *         ->addGroupBy('u.createdAt')
     * </code>
     *
     * @param mixed $groupBy The grouping expression.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function addGroupBy($groupBy)
    {
        if (empty($groupBy)) {
            return $this;
        }

        $groupBy = is_array($groupBy) ? $groupBy : func_get_args();

        return $this->add('groupBy', $groupBy, true);
    }

    /**
     * Specifies a restriction over the groups of the query.
     * Replaces any previous having restrictions, if any.
     *
     * @param mixed $having The restriction over the groups.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function having($having)
    {
        if ( ! (func_num_args() == 1 && $having instanceof CompositeExpression)) {
            $having = new CompositeExpression(CompositeExpression::TYPE_AND, func_get_args());
        }

        return $this->add('having', $having);
    }

    /**
     * Adds a restriction over the groups of the query, forming a logical
     * conjunction with any existing having restrictions.
     *
     * @param mixed $having The restriction to append.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function andHaving($having)
    {
        $having = $this->getQueryPart('having');
        $args = func_get_args();

        if ($having instanceof CompositeExpression && $having->getType() === CompositeExpression::TYPE_AND) {
            $having->addMultiple($args);
        } else {
            array_unshift($args, $having);
            $having = new CompositeExpression(CompositeExpression::TYPE_AND, $args);
        }

        return $this->add('having', $having);
    }

    /**
     * Adds a restriction over the groups of the query, forming a logical
     * disjunction with any existing having restrictions.
     *
     * @param mixed $having The restriction to add.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function orHaving($having)
    {
        $having = $this->getQueryPart('having');
        $args = func_get_args();

        if ($having instanceof CompositeExpression && $having->getType() === CompositeExpression::TYPE_OR) {
            $having->addMultiple($args);
        } else {
            array_unshift($args, $having);
            $having = new CompositeExpression(CompositeExpression::TYPE_OR, $args);
        }

        return $this->add('having', $having);
    }

    /**
     * Specifies an ordering for the query results.
     * Replaces any previously specified orderings, if any.
     *
     * @param string $sort The ordering expression.
     * @param string $order The ordering direction.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function orderBy($sort, $order = null)
    {
        return $this->add('orderBy', $sort . ' ' . (! $order ? 'ASC' : $order), false);
    }

    /**
     * Adds an ordering to the query results.
     *
     * @param string $sort The ordering expression.
     * @param string $order The ordering direction.
     * @return QueryBuilder This QueryBuilder instance.
     */
    public function addOrderBy($sort, $order = null)
    {
        return $this->add('orderBy', $sort . ' ' . (! $order ? 'ASC' : $order), true);
    }

    /**
     * Get a query part by its name.
     *
     * @param string $queryPartName
     * @return mixed $queryPart
     */
    public function getQueryPart($queryPartName)
    {
        return $this->sqlParts[$queryPartName];
    }

    /**
     * Get all query parts.
     * @return array $sqlParts
     */
    public function getQueryParts() { return $this->sqlParts; }

    /**
     * Reset SQL parts
     *
     * @param array $queryPartNames
     * @return QueryBuilder
     */
    public function resetQueryParts($queryPartNames = null)
    {
        if (is_null($queryPartNames)) {
            $queryPartNames = array_keys($this->sqlParts);
        }

        foreach ($queryPartNames as $queryPartName) {
            $this->resetQueryPart($queryPartName);
        }

        return $this;
    }

    /**
     * Reset single SQL part
     *
     * @param string $queryPartName
     * @return QueryBuilder
     */
    public function resetQueryPart($queryPartName)
    {
        $this->sqlParts[$queryPartName] = is_array($this->sqlParts[$queryPartName])
            ? array() : null;

        $this->state = self::STATE_DIRTY;

        return $this;
    }

    /**
     * Converts this instance into a SELECT string in SQL.
     *
     * @return string
     */
    private function getSQLForSelect()
    {
        $query = 'SELECT ' . implode(', ', $this->sqlParts['select']) . ' FROM ';

        $fromClauses = array();

        // Loop through all FROM clauses
        foreach ($this->sqlParts['from'] as $from) {
            $fromClause = $from['table'] . ' ' . $from['alias'];

            if (isset($this->sqlParts['join'][$from['alias']])) {
                foreach ($this->sqlParts['join'][$from['alias']] as $join) {
                    $fromClause .= ' ' . strtoupper($join['joinType'])
                                 . ' JOIN ' . $join['joinTable'] . ' ' . $join['joinAlias']
                                 . ' ON ' . ((string) $join['joinCondition']);
                }
            }

            $fromClauses[$from['alias']] = $fromClause;
        }

        // loop through all JOIN clasues for validation purpose
        foreach ($this->sqlParts['join'] as $fromAlias => $joins) {
            if ( ! isset($fromClauses[$fromAlias]) ) {
                throw QueryException::unknownFromAlias($fromAlias, array_keys($fromClauses));
            }
        }

        $query .= implode(', ', $fromClauses)
                . ($this->sqlParts['where'] !== null ? ' WHERE ' . ((string) $this->sqlParts['where']) : '')
                . ($this->sqlParts['groupBy'] ? ' GROUP BY ' . implode(', ', $this->sqlParts['groupBy']) : '')
                . ($this->sqlParts['having'] !== null ? ' HAVING ' . ((string) $this->sqlParts['having']) : '')
                . ($this->sqlParts['orderBy'] ? ' ORDER BY ' . implode(', ', $this->sqlParts['orderBy']) : '');

        return ($this->maxResults === null && $this->firstResult == null)
            ? $query
            : $this->connection->getDatabasePlatform()->modifyLimitQuery($query, $this->maxResults, $this->firstResult);
    }

    /**
     * Converts this instance into an UPDATE string in SQL.
     *
     * @return string
     */
    private function getSQLForUpdate()
    {
        $table = $this->sqlParts['from']['table'] . ($this->sqlParts['from']['alias'] ? ' ' . $this->sqlParts['from']['alias'] : '');
        $query = 'UPDATE ' . $table
               . ' SET ' . implode(", ", $this->sqlParts['set'])
               . ($this->sqlParts['where'] !== null ? ' WHERE ' . ((string) $this->sqlParts['where']) : '');

        return $query;
    }

    /**
     * Converts this instance into a DELETE string in SQL.
     *
     * @return string
     */
    private function getSQLForDelete()
    {
        $table = $this->sqlParts['from']['table'] . ($this->sqlParts['from']['alias'] ? ' ' . $this->sqlParts['from']['alias'] : '');
        $query = 'DELETE FROM ' . $table . ($this->sqlParts['where'] !== null ? ' WHERE ' . ((string) $this->sqlParts['where']) : '');

        return $query;
    }

    /**
     * Gets a string representation of this QueryBuilder which corresponds to
     * the final SQL query being constructed.
     *
     * @return string The string representation of this QueryBuilder.
     */
    public function __toString() { return $this->getSQL(); }

    /**
     * Create a new named parameter and bind the value $value to it.
     *
     * This method provides a shortcut for PDOStatement::bindValue
     * when using prepared statements.
     *
     * The parameter $value specifies the value that you want to bind. If
     * $placeholder is not provided bindValue() will automatically create a
     * placeholder for you. An automatic placeholder will be of the name
     * ':dcValue1', ':dcValue2' etc.
     *
     * For more information see {@link http://php.net/pdostatement-bindparam}
     *
     * Example:
     * <code>
     * $value = 2;
     * $q->eq( 'id', $q->bindValue( $value ) );
     * $stmt = $q->executeQuery(); // executed with 'id = 2'
     * </code>
     *
     * @license New BSD License
     * @link http://www.zetacomponents.org
     * @param mixed $value
     * @param mixed $type
     * @param string $placeHolder the name to bind with. The string must start with a colon ':'.
     * @return string the placeholder name used.
     */
    public function createNamedParameter( $value, $type = \PDO::PARAM_STR, $placeHolder = null )
    {
        if ( $placeHolder === null ) {
            $this->boundCounter++;
            $placeHolder = ":dcValue" . $this->boundCounter;
        }
        $this->setParameter(substr($placeHolder, 1), $value, $type);

        return $placeHolder;
    }

    /**
     * Create a new positional parameter and bind the given value to it.
     *
     * Attention: If you are using positional parameters with the query builder you have
     * to be very careful to bind all parameters in the order they appear in the SQL
     * statement , otherwise they get bound in the wrong order which can lead to serious
     * bugs in your code.
     *
     * Example:
     * <code>
     *  $qb = $conn->createQueryBuilder();
     *  $qb->select('u.*')
     *     ->from('users', 'u')
     *     ->where('u.username = ' . $qb->createPositionalParameter('Foo', PDO::PARAM_STR))
     *     ->orWhere('u.username = ' . $qb->createPositionalParameter('Bar', PDO::PARAM_STR))
     * </code>
     *
     * @param  mixed $value
     * @param  mixed $type
     * @return string
     */
    public function createPositionalParameter($value, $type = \PDO::PARAM_STR)
    {
        $this->boundCounter++;
        $this->setParameter($this->boundCounter, $value, $type);
        return "?";
    }
}

class Doctrine\DBAL\Statement implements \IteratorAggregate, Doctrine\DBAL\Driver\Statement
{
    protected $sql;    /** @var string The SQL statement. */
    protected $params = array();  /** @var array The bound parameters. */
    protected $stmt;   /** @var Doctrine\DBAL\Driver\Statement The underlying driver statement. */
    protected $platform;  /** @var Doctrine\DBAL\Platforms\AbstractPlatform The underlying database platform. */
    protected $conn;   /** @var Doctrine\DBAL\Connection The connection this statement is bound to and executed on. */

    /**
     * Creates a new <tt>Statement</tt> for the given SQL and <tt>Connection</tt>.
     *
     * @param string $sql The SQL of the statement.
     * @param Doctrine\DBAL\Connection The connection on which the statement should be executed.
     */
    public function __construct($sql, Connection $conn)
    {
        $this->sql = $sql;
        $this->stmt = $conn->getWrappedConnection()->prepare($sql);
        $this->conn = $conn;
        $this->platform = $conn->getDatabasePlatform();
    }

    /**
     * Binds a parameter value to the statement.
     *
     * The value can optionally be bound with a PDO binding type or a DBAL mapping type.
     * If bound with a DBAL mapping type, the binding type is derived from the mapping
     * type and the value undergoes the conversion routines of the mapping type before
     * being bound.
     *
     * @param $name The name or position of the parameter.
     * @param $value The value of the parameter.
     * @param mixed $type Either a PDO binding type or a DBAL mapping type name or instance.
     * @return boolean TRUE on success, FALSE on failure.
     */
    public function bindValue($name, $value, $type = null)
    {
        $this->params[$name] = $value;
        if ($type !== null) {
            if (is_string($type)) {
                $type = Type::getType($type);
            }
            if ($type instanceof Type) {
                $value = $type->convertToDatabaseValue($value, $this->platform);
                $bindingType = $type->getBindingType();
            } else {
                $bindingType = $type; // PDO::PARAM_* constants
            }
            return $this->stmt->bindValue($name, $value, $bindingType);
        } else {
            return $this->stmt->bindValue($name, $value);
        }
    }

    /**
     * Binds a parameter to a value by reference.
     *
     * Binding a parameter by reference does not support DBAL mapping types.
     *
     * @param string $name The name or position of the parameter.
     * @param mixed $value The reference to the variable to bind
     * @param integer $type The PDO binding type.
     * @return boolean TRUE on success, FALSE on failure.
     */
    public function bindParam($name, &$var, $type = PDO::PARAM_STR)
    {
        return $this->stmt->bindParam($name, $var, $type);
    }

    /**
     * Executes the statement with the currently bound parameters.
     *
     * @return boolean TRUE on success, FALSE on failure.
     */
    public function execute($params = null)
    {
        $hasLogger = $this->conn->getConfiguration()->getSQLLogger();
        if ($hasLogger) {
            $this->conn->getConfiguration()->getSQLLogger()->startQuery($this->sql, $this->params);
        }

        $stmt = $this->stmt->execute($params);

        if ($hasLogger) {
            $this->conn->getConfiguration()->getSQLLogger()->stopQuery();
        }
        $this->params = array();
        return $stmt;
    }

    /**
     * Closes the cursor, freeing the database resources used by this statement.
     * @return boolean TRUE on success, FALSE on failure.
     */
    public function closeCursor() { return $this->stmt->closeCursor(); }

    /** Returns the number of columns in the result set. */
    public function columnCount() { return $this->stmt->columnCount(); }

    /**
     * Fetches the SQLSTATE associated with the last operation on the statement.
     * @return string
     */
    public function errorCode() { return $this->stmt->errorCode(); }

    /**
     * Fetches extended error information associated with the last operation on the statement.
     * @return array
     */
    public function errorInfo() { return $this->stmt->errorInfo(); }

    public function setFetchMode($fetchStyle, $arg2 = null, $arg3 = null)
    {
        return $this->stmt->setFetchMode($fetchStyle, $arg2, $arg3);
    }

    public function getIterator() { return $this->stmt; }

    /**
     * Fetches the next row from a result set.
     *
     * @param integer $fetchStyle
     * @return mixed The return value of this function on success depends on the fetch type.
     *               In all cases, FALSE is returned on failure.
     */
    public function fetch($fetchStyle = PDO::FETCH_BOTH)
    {
        return $this->stmt->fetch($fetchStyle);
    }

    /**
     * Returns an array containing all of the result set rows.
     *
     * @param integer $fetchStyle
     * @param mixed $fetchArgument
     * @return array An array containing all of the remaining rows in the result set.
     */
    public function fetchAll($fetchStyle = PDO::FETCH_BOTH, $fetchArgument = 0)
    {
        if ($fetchArgument !== 0) {
            return $this->stmt->fetchAll($fetchStyle, $fetchArgument);
        }
        return $this->stmt->fetchAll($fetchStyle);
    }

    /**
     * Returns a single column from the next row of a result set.
     *
     * @param integer $columnIndex
     * @return mixed A single column from the next row of a result set or FALSE if there are no more rows.
     */
    public function fetchColumn($columnIndex = 0)
    {
        return $this->stmt->fetchColumn($columnIndex);
    }

    /**
     * Returns the number of rows affected by the last execution of this statement.
     * @return integer The number of affected rows.
     */
    public function rowCount() { return $this->stmt->rowCount(); }

    /**
     * Gets the wrapped driver statement.
     * @return Doctrine\DBAL\Driver\Statement
     */
    public function getWrappedStatement() { return $this->stmt; }
}
