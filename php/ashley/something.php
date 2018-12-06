$ f Ashley/Chat
common-new/src/Ashley/ServiceProvider/ChatServiceProvider.php

common-new/src/Ashley/Chat/XmppSettingsProvider.php
common-new/src/Ashley/Chat/FiltersHandler.php // FilterCollection
common-new/src/Ashley/Chat/ChatProvider.php

common-new/src/Ashley/Chat/Filters/FilterInterface.php
    common-new/src/Ashley/Chat/Filters/Country.php
    common-new/src/Ashley/Chat/Filters/Db.php
    common-new/src/Ashley/Chat/Filters/Language.php
    common-new/src/Ashley/Chat/Filters/State.php
    common-new/src/Ashley/Chat/Filters/UserAgent.php

common-new/tests/unit/Chat/ChatProviderTest.php

<?php
namespace Ashley\Chat;

use Ashley\Account\Account;
use Ashley\Main\Entity\SomethingEntity;
use Doctrine\ORM\EntityManager;

/**
 * SomethingProvider class
 */
class SomethingProvider
{
    const CLASS_NAME = __CLASS__;

    /**
     * @var EntityManager
     */
    private $entityManager;

    public function __construct(EntityManager $entityManager)
    {
        $this->entityManager = $entityManager;
    }

    /**
     * @return \Doctrine\ORM\EntityManager
     */
    private function getEntityManager()
    {
        return $this->entityManager;
    }

    /**
     * @param Account $account
     * @return object
     */
    public function fetchSettings(Account $account)
    {
        return $this->getEntityManager()->getRepository(SomethingEntity::CLASS_NAME)
            ->findOneBy(array('pnum' => $account->getPnum()));
    }
}




# common-new/src/Ashley/ServiceProvider/ChatServiceProvider.php

<?php
namespace Ashley\ServiceProvider;

use Ashley\Chat\ChatProvider;
use Ashley\Chat\Filters\Db;
use Ashley\Chat\Filters\Language;
use Ashley\Chat\Filters\State;
use Ashley\Chat\Filters\UserAgent;
use Ashley\Chat\Filters\Country;
use Ashley\Chat\FiltersHandler;
use Ashley\Chat\XmppSettingsProvider;

/**
 * Provide chat service
 */
class ChatServiceProvider implements ServiceProviderInterface
{
    public static function register(\Pimple $app, array $options = array())
    {
        $app['ashley.chat.xmppSettingsProvider'] = function () use ($app) {
            return new XmppSettingsProvider($app['orm']['main']);
        };

        $app['ashley.chat.filtersHandler'] = $app->share(function () {
            return new FiltersHandler();
        });

        $app['ashley.chat.filtersHandler'] = $app->extend(
            'ashley.chat.filtersHandler',
            function ($filtersHandler) use ($app) {
                $filtersHandler->addFilter(new Country($app['ashley.country.configProvider']));
                $filtersHandler->addFilter(new State($app['ashley.country.configProvider']));
                $filtersHandler->addFilter(new Language($app['ashley.lang.configProvider']));
                $filtersHandler->addFilter(new Db($app['ashley.chat.xmppSettingsProvider']));
                $filtersHandler->addFilter(new UserAgent($app['ashley.config']));

                return $filtersHandler;
            }
        );

        $app['ashley.chat'] = function () use ($app) {
            return new ChatProvider($app['ashley.chat.filtersHandler']);
        };
    }
}
