<?php // ���ڵ�Ashley������ڵ����⣬�Ľ�˼·�������뷨�Ƿ������⣿

class AccountServiceProvider implements ServiceProviderInterface
{
    public static function register(\Pimple $app, array $options = array())
    {
        $app['ashley.account'] = $app->share(
            function ($app) {
                return new AccountProvider( // Provider�Ĺ��캯������̫��Ĳ���
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

    public function __construct( // Provider�Ĺ��캯������̫��Ĳ���
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
// �ܷ�ÿ��Provider�Ĺ��캯�����ĳ�ֻ��Ҫһ��$app����

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
    use DependencyTrait;  // Provider��Ҫ��Dependency

    protected $app;

    public function __construct($app)
    {
        $this->app = $app;  // Provider�Ĺ��캯�����ĳ�ֻ��Ҫһ��$app����
    }
}

trait DependencyTrait  // ����ϵͳ�����е�Dependency
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

trait DependencyMock  // ����ϵͳ�����е�Dependency��Mock������дUnitTest
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
