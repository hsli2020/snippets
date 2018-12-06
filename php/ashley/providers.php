<?php

class Ashley\Account\AccountProvider
{
    public function find($pnum)
    public function save(Account $account)

    public static function generateLoginKey($username, $password, $email)

    public function getAddressInfo($pnum)

    public function softDelete(Account $account)
    public function reinstate(Account $account, $unsuspend = true)

    public function setGeoLocation(Account $account, AshleyLocation $data)

    public function hasPurchased(Account $account)

    public function isAccountSince(Account $account, $stringTimeAgo)

    protected function convertToTimestamp($strTime)
}

class Ashley\Account\FavoriteProvider
{
    public function isFavorite($pnum, $targetPnum)
    public function areFavorites($pnum, array $pnums)

    public function getFavoriteCount(Account $account)

    public function favorite(Account $account, Account $favoriteAccount)
    public function unfavorite(Account $account, Account $favoritedAccount)
    public function unfavoriteMulti(Account $account, array $favoritedAccounts)

    public function getFavorites(Account $account, $limit = 20, $offset = 0)
    public function getFavorited(Account $account, $limit = 20, $offset = 0)
    
    public function save(Favorite $favorite, $flush = true)
   
    private function mapAreFavoritesToPnum($pnums, $favorites)
}

class Ashley\Account\BlockProvider
{
    public function isBlocking(Account $account, Account $blocked)
    public function isBlockingOrBlockedBy(Account $account, Account $account2)
    public function getBlockCount(Account $account)

    public function block(Account $blocker, Account $blocked)
    public function unblock(Account $blocker, Account $blocked)
    public function unblockMulti(/*Acccount*/ $blocker, array $accounts)

    public function getBlocking(Account $blocker, $limit = 20, $offset = 0)
    public function getBlockedBy(Account $blocked, $limit = 20, $offset = 0)
    public function getBlockedByPnums(Account $account, $setLimit = false)
    public function getBlockingPnums(Account $account, $setLimit = false)
    public function getBlockedAndBlockerPnums(Account $account, $setLimit = false)
}

class Ashley\Account\FraudProvider
{
    public function find(Account $account)

    public function getLinkedByEmailPnums(Account $account, $email, array $statuses, $limit)

    public function haveFraudPnums(array $pnums, array $linkedTo)

    public function markFraud(Account $account)
    public function markFraudNotify(Account $account)

    public function isFraud(Account $account)

    public function fraudStatusConfirmed(Account $account)
    public function fraudStatusSuspicious(Account $account)
    public function fraudStatusOk(Account $account)
    public function fraudStatusReview(Account $account)
    public function fraudStatusNew(Account $account)

    private function checkFraudStatus($account, $status)

    public function saveFraudStatus(Account $account, FraudStatus $fraudStatus)
}

class Ashley\Message\MessageProvider
{
    public function find ($messageId)

    public function send(Account $from, Account $to, $body, 
                        $key=false, $priority=false, $replyId=0, $isNew=true) 
    public function sendCollect(Account $from, Account $to, $body, $key = false)
    public function sendTravelingMessage(Account $from, Account $to, $body, $key = false)
    public function sendPriorityConfirmation(Account $from, Account $to)
    public function sendMessage(AbstractMessage $message)

    protected function setReplied($messageId, AmMailMessage $replyMessage)

    public function findLastUnreadPriorityMessage(Account $to, Account $from)

    public function markAsRead(Account $account, AmMailMessage $message)
    public function markAsSaved(Account $account, AmMailMessage $message)
    public function markAsPaid(Account $account, AmMailMessage $message, Account $from)
    public function markAsDeleted(Account $account, AmMailMessage $message)

    public function deleteMessagesByIds(array $messageIds, $toPnum = null, $fromPnum = null)
    public function deleteSentMessagesByIds(array $messageIds)
    public function deleteMessageThread($pnum, $withPnum)

    public  function getWinkAndFavMessageSentCount(Account $from, $hour = false)
    private function getFormattedCollectMessage($subject, $body)
}

class Ashley\Wink\WinkProvider
{
    public function wink(Account $from, Account $to, $winkId, $key = false)

    public function customizeWink(Account $from, Account $to, $subject, $body)

    public function getWinkSubject($winkId)
    public function getWinkSubjects(Account $account)
    public function getWinkMap(Account $account)
    public function getWinkBody($winkId, Account $account)
    public function getWinks(Account $account, Translator $translator = null)
}

class Ashley\Gift\GiftProvider
{
    protected function configDefaults()

    public function send(Account $from, Account $to, $giftId, $body)

    public function getGiftItems(Account $from)
    public function getGiftItem(Account $from, $giftId)

    public function getGiftByGiftId($giftId)
    public function getGiftImageUrl($giftImageName, $scheme = 'https')

    public function shouldRateLimit(Account $from)

    protected function validateMessageBody($from, $body, $giftItem)
}

class Ashley\Account\RelationshipProvider
{
    public function isCommunicationOpen(Account $account1, Account $account2)

    protected function isChannelOpen(Account $from, Account $to)

    public function openCommunication(Account $from, Account $to)

    public function isInitialContact(Account $me, Account $other)

    private function getFirstMessageId(Account $from, Account $to)

    public function isPriorityFree(Account $from, Account $to, 
                                   PaymentProvider $paymentProvider, 
                                   CreditProvider $creditProvider) 
    public function isValidSeeking(Account $from, Account $to)
    public function isStraight(Account $account)
    public function isSeekingFemale(Account $account)
    public function isMaleSeekingFemale(Account $account)
    public function isSeekingMale(Account $account)

    public static function getMarriedStatus($seeking)
}



class Ashley\Account\MailProvider
{
    public function getActiveMessages(Account $account, Account $withAccount, $offset = 0, $limit = 10)

    public function getPrivateShowcasePhotos(Account $account, $offset = 0, $limit = self::DEFAULT_SHOWCASE_PHOTOS_LIMIT) 

    private function addGiftDetails($messages = array())

    private function addWinkDetails($messages = array())

    public function markMessagesAsRead($messages, $pnum)

    public function getUnreadMessageCount(Account $account)
}

class Ashley\Account\NoticesProvider
{
    public function detectNotices(Account $account, $type)
    public function markNoticeViewed(Account $account, $type)
}

class Ashley\Account\RegistrationProvider
{
    public function createAccount(array $values, AffiliateTrack $affiliateTrack)

    private function before(array &$values, AffiliateTrack $affiliateTrack)
    private function after(array $values, $pnum)

    private function saveToAminno(array $values)
    private function saveToAm(array $values, $pnum)

    private function operationForInterfed(array $values, $pnum, array $memberEmailData)
    private function execMoreQueries(array $values, $pnum)

    private function getRequestInfo()
    private function getGeoInfo(array $values)

    private function getAffiliateData($values, AffiliateTrack $affiliateTrack)
    private function needToClearAffiliateData($values)

    private function saveToAminnoQueChange($pnum, $fieldname, $valueAfter)
}

class Ashley\Account\RuleProvider
{
    public function get($ruleName)
    public function appliesTo(Account $account, $ruleName)
}

class Ashley\Account\ViewedMeProvider
{
    public function bind(\IteratorAggregate $parameters)

    public function getViewedMePnums($pnum, $pnumsToExclude = array())

    private function applyRefineFilters(QueryBuilder $qb)

    private function getViewedTime($createdon)

    public function viewProfile(Account $from, Account $to)

    private function filterPnums($pnums, $limit)
    private function filter($justPnums)
}

class Ashley\Credit\CreditProvider
{
    public function getEngagerCreditBonus(Account $account)
    public function getEngagerCreditBonusByTransId(Account $account, $id)

    protected function decodeBonusInfo($comment)

    public function getCreditCount(Account $account)

    public function spend(Account $account, CreditCost $cost)

    protected function processLine(array $line)

    public function addCredits (Account $account, $credits, $itemCode)
}

class Ashley\Country\CountryConfigProvider
{
    public function fetchByCountryCode($countryCode)
    public function fetchByCountryId($countryId)

    protected function isEnabled($countryId)

    public function codeToId($code)
    public function idToCode($id)

    public function codeToCountryName($code)

    protected function getConfig()
    protected function getConfigByCountry($id)

    public function isInterfedCountry($countryId)
    public function isExportEnabledInterfedCountry($countryId)
    public function isBCMsgEnabledInterfedCountry($countryId)

    public function getCountryIdsUsingCity()
    public function getCountryIdsWithoutCityZip()

    public function getAddressFormat($countryId)
    public function getSearchAddressFormat($countryId)

    public function getContactPhoneNumber(Account $user)
    public function getOfficeHours($countryId)
    public function getOperatingStatus($countryCode)
}

class Ashley\Conversation\ConversationProvider
{
    private function mapMailboxes(array $mailboxes)

    private function loadConversations(Account $account, $mailBox)

    public function getConversations(Account $account, $mailBox, $offset = 0, $limit = 10)
    public function countConversations(Account $account, $mailBox)

    private function messagesToConversations($pnum, array $messages)
    private function sortMessages(array $messages)

    private function filterByMailbox($mailBox, array $conversations)
}

class Ashley\Country\StateProvider
{
    public function find($id)
    public function findByStateName($value)
}

class Ashley\Email\PromotionProvider
{
    public function getPromos($countryId = 0)
    public function getPromotions()

    public function insertPromotion($lastInsertId, $params)
    public function updateCheckPromo($promoId)
    public function addPromotion($params)

    public function updatePromotion($params)

    public function getPromoTimes($countryId, $email_id)
    public function getPromoKey($countryId = 0)

    public function getPromotionDetails($promoId)

    public function emailIdAlreadyUsed($emailId, $promoId)
}

class Ashley\Email\EmailProvider
{
    public function queue($type, Account $account, array $args = [])
    public function queueRunnerAt($at, $templateName, Account $account, $queueName, $jobClass, array $args = [])
    public function queueRunner($templateName, Account $account, $queueName, $jobClass, $isECD, array $args = [])

    public function getJobName($type)
    public function getJobClass($type)

    public function shouldReceiveEmail(Account $account)

    private function isValidEmail(Account $account)
}

class Ashley\Email\EmailPromotionProvider
{
    public function getNextPromotionId()

    public function getCSVName($email_id)

    public function removePnumsFromQueueTable($deletePnums)

    public function getPnumsToDelete($pnum_limit)
    public function getPnumsFromQueueTable($limit = self::GET_NUMBER_OF_PNUMS)
    public function getPnumsFromFile($csv, $offset)

    public function flagAsLoaded($email_id)

    public function queuePnums($pnums, $queueNow = false)

    public function requeueMissedMessages()
    public function selectRequeueMessages()

    public function importPnumList($email_id, $pnums)

    public function getVolumeFromCSV($csv)
    public function getCSVfileList()
}

class Ashley\Package\PackageProvider
{
    private function loadPinfPackages($packageId = null)
    private function fillPinfPackage($packageId, $package)

    public function getPackage($packageId)
    public function getPackages($isRequestMobile = false, $filterIds = null, $sortByPrice = true)

    public function hasSubscription($packageId)

    private function sortPackagesByPrice(array $packages)
    private function exludeNotMobilePackages(array $packages)

    public function getPackageCurrency()
    public function getDisclaimers(array $ids)
}

class Ashley\Payment\DatabaseExchangeRateProvider implements ExchangeRateProvider
{
    public function getAllRates()
    public function getAllNewRates()

    private function loadRates()
    private function refreshRates()
}

class Ashley\Pricing\CurrencyProvider
{
    public function getDetailsFromId($currencyId)
    public function getDetailsFromCode($code)

    public function convertIdToCode($currencyId)

    public function createFromId($currencyId)
    protected function createFromEntity(CurrencyEntity $currency)
    public function createFromCode($code)

    public function getDisplayCurrencyForCountryId($countryId)
    public function getChargeCurrencyForCountryId($countryId)
    private function getCountryCurrencyForCountryId($countryId)
}

class Ashley\Pimg\PimgProvider
{
    public function storePhoto($pnum, $photoId, $photoFilepath)
    public function cropAndStorePhoto($pnum, $sourcePhotoId, $sourcePhotoSize, $croppedPhotoId, $cropDetails)

    public function createPhotoFromSVG($pnum, $sourcePhotoId, $sourcePhotoSize, $newPhotoId, $svg)

    public function getMaskDetails($pnum, $sourcePhotoId, $sourcePhotoSize, $displayWidth)
    public function getFaces($pnum, $photoId)

    public function purgePhotos($pnum, $photoIds)

    public static function getPhotoSizes()
    public static function getImageSizeCode($imageSize)

    public function getPhotoUrls($pnum, $photoId, $scheme = 'https')
    public function getPhotoUrlsBySize($pnum, $photoId, $photoSizes, $scheme = 'https')
    public function getPhotoUrl($pnum, $photoId, $size = 'normal', $scheme = 'https')

    private function getUrlHash($pnum, $imageId, $imageSize)
    private function getImagePurgeHash($pnum, $imageIds)

    public function getCDNUrl()
}

class Ashley\Payment\PaymentProvider
{
    public function hasFemaleInitiatedContact(Account $account)
    public function getLastApprovedTransaction(Account $account)
}

class Ashley\Photo\MasksProvider
{
    public function getMasks($gender = 'all', $enabledOnly = true)
    public function getBlurs()

    private function filterDisabled($masksConfig, $enabledOnly)

    private function translate($text)
}

class Ashley\Photo\PhotoProvider
{
    public function fetchAllByUser(UserInterface $user)
    public function fetchByIdAndUser($photoId, UserInterface $user, $activeOnly = true)

    public function createPhoto(UserInterface $user, $display, $imagePathname)
    public function createPhotoFromUrl(UserInterface $user, $display, $photoUrl)
    public function createPhotoFromSVG(UserInterface $user, $sourcePhotoId, $size, $svg)
    public function createPhotoFromCrop(UserInterface $user, $sourcePhotoId, $size, $cropDetails)

    public function updatePhotoDisplay(UserInterface $user, $photoId, $display)

    public function deletePhoto(UserInterface $user, $photoId)
    public function deletePhotos(UserInterface $user, array $photoIds)

    public function getPurgingPhotos(UserInterface $user)
    public function purgePhotos(UserInterface $user, array $photoIds)

    public static function getSupportedImageTypes()

    public function storePhoto($pnum, $file, $display, $isLarge)
    public function updatePhotoApproval(UserInterface $user, $photoId, $approved, $setFeaturedIfNone = false)

    private function queueFirstApproval($pnum, $photoId, $display, $createdon)
    private function queueUpdate(Photo $photo)

    private function clearCache($pnum)

    public function getActivePhotos($pnum)
    public function getActivePhotosCount($pnum)
    public function getPhotosByPnums($pnums, $approvedAndNotHiddenOnly = true, $activeOnly = true)
    public function getPhotoUrls($pnums, $photoSizes, $approvedAndNotHiddenOnly = true, $activeOnly = true)
    public function getFeaturedPhotoUrls($pnums, $photoSize)
    public function getFeaturedPhotoUrl($pnum, $photoSize)

    public function hasPublicPrivatePhotos($pnum)

    public static function getMaxUploadSize()

    public function storePhotoFacesDetail($pnum, $photoId)
}

class Ashley\Photo\PhotoSourceProvider
{
    public function logCropPhoto($newPhotoId, $sourceId, $displayWidth, $cropDetails)
    public function logMaskPhoto($newPhotoId, $sourceId, $maskDetails)

    public function getCropData($photo_id)
    public function getMaskData($photo_id)
    public function getSourcePhotoId($photo_id)

    private function findSourceId($photo_id)
}

class Ashley\Fraud\SuspiciousIpProvider
{
    public function isSuspicious($ipnum)

    public function mark($ipnum)
    public function clear($ipnum)
    public function toggle($ipnum)
    
    private function find($ipnum)
}

class Ashley\Feature\FemaleInitiatedContactProvider
{
    public function isCriteriaMet(Account $from, Account $to)

    public function openCommunicationChannel(Account $from, Account $to)

    private function getRelationshipProvider()
    private function getPaymentProvider()
}

class Ashley\Profile\ProfileProvider
{
    public function getProfile(Account $account)

    private function addProfileAttributes($profile, Account $account)

    public function getProfilesByPnums(Account $account, $pnums, $type = 'short', $activeOnly = true, $associativeArray = false) 

    private function fetchProfileDetails($pnums, Account $account, $type, $activeOnly = true)
    private function getFullProfileFields()
    private function formatProfileValues(&$profileDetails, $type, Account $account)

    private function isPaidContact($profileDetail)
    private function isProfileAvailable($profileDetail)

    private function addPhotosToProfiles(&$profileDetails, $pnums, Account $account)
    private function addMessageCostAndTranslate(Account $account, &$profileDetails, $pnums, $type)
    private function addPhoneCallToProfile(&$profile, Account $fromAccount)

    public function getCommChannelOpen(Account $from, Account $to)
}

class Ashley\QuickReply\QuickReplyMessageProvider
{
    public function getRandomMessageList(Account $account, $amount = QuickReplyUtils::DEFAULT_RANDOM_MESSAGES)
    public function getRandomMessage(Account $account, $amount = QuickReplyUtils::DEFAULT_RANDOM_MESSAGES)
    public function getAllMessages()

    private function getQuickReplyConfig()

    public function getPredefinedMessage($predfinedMessageId)
    public function getOldUnreadMessageCount($days, Account $account)
    public function getUnreadMessageCount(Account $account)
    public function getUnreadMessages(Account $account, $callFromAdmin)
    public function getUnreadMessagesFrom(Account $sender, Account $receiver)
    public function markMessageAsReadReplied($messageId, $replied)
    public function sentRecently($userFromPnum, $userToPnum, $days)
    public function getQuickReplyCount($userto)
    public function receivedRecently($userto)
    public function getTranslatedMessage($selectedReplyOptionId, $languageId)

    private function getMessageTranslationKey($predfinedMessageId)
}

class Ashley\QuickReply\QuickReplyProvider
{
    public function isEligibleQuickReply(Account $account)

    public function quickReplyDisabled(Account $account)

    public function isEnabled(Account $account)

    public function getQuickReplySetting(Account $account)

    public function saveQuickReplySetting($data, Account $account, $callFromAdmin = null)

    public function disableQuickReply(Account $account)

    public function needShowPopup(Account $account)
}

class Ashley\Research\ResearchProvider
{
    public function queue($type, $data, Account $account, Account $otherAccount = null)

    private function loggable($type, $data, $account, $otherAccount)
    public function canLog(Account $account)

    private function jobEnabled($type)
    public function enabled()

    public function checkProfiles($job, $currentPnum, $otherPnum = false)
    public function checkProfileExist($pnum)

    private function queueProfileInsert($pnum)
}

class Ashley\Resource\OptionsProvider
{
    private function setCountryConfig($countryConfigOrId = null)

    public function get($optionName, $countryConfigOrId = null)
    public function getOptionById($optionName, $optionId, $countryConfigOrId = null)

    private function getDateOptions($optionName)
    private function getResourceOptions($optionName)
    private function getVisibleCountries()
    private function getCountryCode()
    private function getUnitSystem()
    private function getDecimalPoint()
}

class Ashley\MemberActivityLevel\MemberActivityLevelProvider
{
    public function getUsers($maxPnum)
    public function getMinAndMaxPnum()

    public function queueUpdateActivityLevel($pnum)
    public function updateActivityLevel(Account $account, $active_photos)

    public function isTrackedAccount(Account $account)

    public function getMemberActivityLevel(Account $account)
    public function addMemberActivityLevel(Account $account)
    public function setMemberActivityLevel(Account $account, $new_level_id)
    public function logMemberActivityLevelHistory(Account $account, $level_id)

    public function hasQuickReplyDisabled(Account $account)
    public function hasCustomContent(Account $account)
    public function isDisengaged(Account $account)
    public function getNextEmailToSend(Account $account, $level_id, $changed_at)
    public function queueEmail(Account $account, $email_type_id)
    public function receivesECDEmails(Account $account)
}

class Ashley\Messenger\MessengerProvider
{
    public function signProfileNumberOutOfMessenger($pnum, $deviceId, $deviceType)
    public function signProfileNumberIntoMessenger($pnum, $deviceId, $deviceType)

    public function isOnline(Account $receiver)

    public function logMessengerPhoto(Account $sender, Account $receiver)
    public function logMessengerChat(Account $sender, Account $receiver, $message)
    public function logMessengerEmoji(Account $sender, Account $receiver)

    public function createNewMessengerConversation(Account $sender, Account $receiver)

    public function buildEmojiResponse(Account $sender, Account $receiver, $emojiCost)
    public function buildMessageResponse(Account $sender, Account $receiver)
    public function buildChatResponse(MessengerChatResponse $chatResponse)

    public function prepareMessageToBeSent(MessengerChatResponse $chatResponse)

    public function sendMessageReceivedPushNotifications(Account $sender, Account $receiver)
    public function removeInvalidDeviceTokens($pnum, $invalidDeviceIds)
}

class Ashley\Messenger\Roster\DoctrineRosterProvider implements RosterProvider
{
    public function findByAccount(Account $account)

    private function addContactsAddedInMessenger(Account $account, Roster $roster)
    private function addFavoritedAccount(Account $account, Roster $roster)
    private function addBlockedAccount(Account $account, $roster)

    public function addSpokenWithAccount(Account $account, Roster $roster)
    public function getSpokenWithProfileDetails(array $conversations)

    public function findAccountsContactedInMessenger(Account $account)
    public function findAccountsWhoContactedInMessenger(Account $account)

    private function getNicknames(array $profileNumbers)
}

class Ashley\Messenger\MessengerRedisProvider implements MessengerSessionsInterface
{
    public function getRedisKeyName($pnum)

    public function getProfileNumberStatus($pnum)
    public function setProfileNumberStatus($pnum, $status)

    public function addDeviceToken($pnum, $deviceId, $deviceType)

    public function removeValidDeviceToken($pnum, $deviceId, $deviceType)
    public function removeDeviceToken($pnum, $deviceId, $deviceType)
    public function removeInvalidDeviceToken($pnum, $deviceId, $deviceType)

    public function getRemainingDeviceTokensOfType($pnum, $deviceId, $deviceType)

    public function getDeviceTokensOfType($pnum, $deviceType)

    public function getHashKeyForDeviceType($deviceType)

    public function doesPnumHaveTokensLeft($pnum)
}

class Ashley\Messenger\MessengerCommunicationProvider
{
    public function purchaseEmoji(Account $sender, Account $receiver, $emojiCost)
    public function purchaseCommunicationChannel(Account $sender, Account $receiver, $costToSendMessage)

    public function spendCredits(Account $from, Account $to, $credits, $type)

    public function doesTheReceiverNeedToBuyCreditsToRespond(Account $receiver, Account $sender)

    public function getAccountCreditCount(Account $account)
    public function getMessageCost(Account $sender, Account $receiver)

    public function assertCanContact(Account $from, Account $to)
}

class Ashley\Profile\ActionsProvider
{
    public function markProfilesAsViewed(Account $fromAccount, $pnums)
    public function markProfileForReview(Account $fromAccount, Account $reportedAccount)
}

class Ashley\Search\SearchProvider
{
    public function search(Searchable $searchForm, $blockedPnums = array(), SearchHistory $searchHistory = null, $addPaidSearchProfiles = true) 

    private function addPaidSearchProfiles($searchedPnums, $searchTerms, $filters, $langId, $label)

    private function mergePaidPnums($pnums, $paidPnums)
}

class Ashley\Showcase\ShowcaseProvider
{
    public function grant(Account $from, Account $to)
    public function grantKey(Account $showcaseRequester, Account $showcaseOwner, $granted = false)

    public function revoke(Account $from, Account $to)
    public function revokeKey(Account $showcaseRequester, Account $showcaseOwner)

    public function request(Account $from, Account $to)

    public function accept(Account $from, Account $to)
    public function decline(Account $from, Account $to)
    public function declineKey(Account $showcaseRequester, Account $showcaseOwner)

    protected function queueJob($jobName, Account $from, Account $to, $messageId = null)

    public function managePendingKeys($pnum, $display)
    public function revokePendingKey(Account $showcaseRequester, Account $showcaseOwner)
    public function addPendingKey(Account $showcaseRequester, Account $showcaseOwner)

    public function hasPendingKey(Account $showcaseRequester, Account $showcaseOwner)
    public function hasKeyOrPendingKey($pnum_requester, $pnum_owner)
    public function getPendingKeys(Account $to)
    public function loadPendingKeySentInformation($pnum)

    private function getKey ($pnum)

    public function keyCreatedDate($pnum)

    public function hasPrivateShowcase($pnum)
    public function hasShowcase($pnum)
    public function hasPendingOnlyShowcase($pnum)
    public function hasPendingShowcase($pnum)
    public function hasPrivateShowcaseAccess(Account $showcaseRequester, Account $showcaseOwner)

    private function getPendingInstance($pnum)
}

class Ashley\Traveling\TravelingProvider
{
    public function sendTravelingMessages(Account $from, array $toPnums, $body, $key = false)

    public function incrementSearchedCount(Account $from)

    public function getKey($from, $key)

    public function getSearchedCount($pnum)

    public function isEnabled($seekingId, $fraudStatus, Account $account)

    public function isSearchLimitReached(Account $account)

    private function logTravelingSearch(Account $from, $transactionData, $toPnums)

    private function setShowTravelingSearch(Account $from, $transactionData)

    public function setTransactionData(Account $account, array $data = array())
    public function getTransactionData(Account $account, $property = '')
    public function clearTransactionData(Account $account)

    public function isLocationTooClose(Account $account, $latitude, $longitude)

    private function getMinimumDistance(Account $account)

    public function isOverlappingSearch(Account $account, $city, \DateTime $startDate, \DateTime $endDate)
}

class Ashley\Terms\TermsProvider
{
    public function isShowTermsForUserFromAffiliate($pnum)

    public function agreeTerms($pnum, $terms)

    public function affiliateUserAgreeTerms($pnum)

    public function isUserFromAffiliate($pnum)
    public function isAffiliateUserSignedTerms($pnum)

    private function getKey($pnum)

    private function getAffliateTermsName()
}

class Avid\Billing\SubscriptionProvider
{
    protected function getSubRepo()

    protected function getSubscription(MemberAccountSubscription $sub)

    public function queueRebills()

    public function find($id)
    public function create($pnum, Package $package)

    public function cancelAllForAccount($pnum)

    public function getForAccount($pnum)
    public function getLatestByAccount($pnum)
    public function getActiveForAccount($pnum, $package)
}

class Ashley\Push\MobileNotificationProvider implements PushNotificationInterface
{
    public function push($fromPnum, $toPnum, $messageType, array $params)
    private function checkDeviceAndPush($type, $message, $deviceToken, $key, $toAccount)
}

class Ashley\Response\CreditCardResponseProvider
{
    public function response(array $pinfResponse, array $packages, array $currency)

    private function hasReponseByStatus($status)

    private function mapFields(array $pinfResponse)
}

class Ashley\Security\OAuth2\Model\RedisClientProvider
{
    public function createClient($clientId, $clientSecret, $redirectUri, $grantType, $scope)

    public function getClient($clientId)

    public function deleteClient($clientId)

    public function checkClientCredentials(OAuth2Client $client, $clientSecret)

    private function getClientIdKey($clientId)
}

abstract class Ashley\Security\OAuth2\Model\RedisTokenProvider
{
    abstract protected function getTokenKey($token);
    abstract protected function getPnumKey($pnum);

    public function createToken($token, $clientId, $pnum, $expires, $scope)

    protected function getTokenData($token)

    public function deleteToken($token, $pnum)
    public function deleteTokens($pnum, $expiresBefore = null)

    public function getTokenCount($pnum)

    protected function removeExpiredTokensFromTokenSet($pnum)
}

class Ashley\Provider\ProfileImageProvider
{
    public function getCachedImagePathByImageID($profileNumber, $imageId, $imageSize)

    public function getCachedImagePathByHash($profileImageHash)
    public function getImageByHash($profileImageHash)

    private function getImage(ProfileImage $profileImage)
    private function getCachedImagePath(ProfileImage $profileImage)

    public function purgeImages($pnum, $imageIds, $hash)
    public function getImagePath($profileNumber, $imageId)

    public function cropAndSaveImage($sourcePath, $sourceOriginalPath, $destinationPath, $cropDetails)
    public function createPhotoFromSVG($sourcePath, $destinationPath, $svg)

    private function deleteImage($path)

    public function getImagePurgeHash($pnum, $imageIds)
    public function getImageSubPath($pnum)

    public function getCachedImageSubPath($pnum, $imageId, $imageSize)
    public function fixImage($pnum, $imageId)
    public function getOriginalImage($profileNumber, $imageId, SymfonyFile $imageData)

    private function guardAgainstOriginalImageNotFound($originalImageFile)
    private function resizeImage($originalImageFile, $cachedImageFile, $width, $height, $quality, $squareCrop) 
    private function getProfileImageByHash($profileImageHash)
}
