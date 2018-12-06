<?php // 现在的Ashley代码存在的问题，改进思路，这种想法是否有问题？

class AccountServiceProvider implements ServiceProviderInterface
{
    public static function register(\Pimple $app, array $options = array())
    {
        $app['ashley.account'] = $app->share(
            function ($app) {
                return new AccountProvider( // Provider的构造函数带了太多的参数
                    $app['orm']['main'],
                    $app['ashley.log'],
                    $app['dbs']['aminno_write'],
                    $app['ashley.resque']['events']
                );
            }
        );
        $app['ashley.account.block'] = $app->share(
            function ($app) { return new BlockProvider($app['orm']['main']); }
        );
        $app['ashley.account.favorite'] = $app->share(
            function ($app) { return new FavoriteProvider($app['orm']['bmdb']); }
        );
    }
}

class AccountProvider
{
    protected $em;
    protected $activityLog;
    protected $dbWrite;
    protected $resque;

    public function __construct( // Provider的构造函数带了太多的参数
        EntityManager $em,
        LogProvider $activityLog,
        Connection $dbWrite,
        ResqueJobProvider $resque
    ) {
        $this->em = $em;
        $this->activityLog = $activityLog;
        $this->dbWrite = $dbWrite;
        $this->resque = $resque;
    }
}
--------------------------------------------------------------------------------
// 能否将每个Provider的构造函数都改成只需要一个$app变量

class AccountServiceProvider implements ServiceProviderInterface
{
    public static function register(\Pimple $app, array $options = array())
    {
        $app['ashley.account'] = $app->share(
            function ($app) {
                return new AccountProvider($app);
            }
        );
        $app['ashley.account.block'] = $app->share(
            function ($app) {
                return new BlockProvider($app);
            }
        );

        $app['ashley.account.favorite'] = $app->share(
            function ($app) {
                return new FavoriteProvider($app);
            }
        );
    }
}

class AccountProvider
{
    use DependencyTrait;  // Provider需要的Dependency

    protected $app;

    public function __construct($app)
    {
        $this->app = $app;  // Provider的构造函数都改成只需要一个$app变量
    }
}

trait DependencyTrait  // 包含系统中所有的Dependency
{
    private function getAccountProivider()
    {
        if (!$this->accountProvider)
            $this->accountProvider = $app['ashley.account'];

        return $this->accountProvider;
    }

    private function getBlockProivider()
    {
        if (!$this->blockProvider)
            $this->blockProvider = $app['ashley.account.block'];

        return $this->blockProvider;
    }

    private function getFavoriteProivider() { return $app['ashley.account.favorite']; }
}

trait DependencyMock  // 包含系统中所有的Dependency的Mock，便于写UnitTest
{
    private function getAccountProivider()
    {
        return $this->getMockBuilder('\Ashley\Account\Account')
                    ->disableOriginalConstructor()->getMock();
    }

    private function getBlockProivider()
    {
        return $this->getMockBuilder('\Ashley\Account\Block')
                    ->disableOriginalConstructor()->getMock();
    }

    private function getFavoriteProivider()
    {
        return $this->getMockBuilder('\Ashley\Account\Favorite')
                    ->disableOriginalConstructor()->getMock();
    }
}
