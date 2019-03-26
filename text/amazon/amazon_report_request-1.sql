-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.38-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             10.1.0.5464
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table bte.amazon_report_request
CREATE TABLE IF NOT EXISTS `amazon_report_request` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `store` varchar(20) NOT NULL,
  `report_type` varchar(80) NOT NULL,
  `filename` varchar(80) NOT NULL,
  `tablename` varchar(80) NOT NULL,
  `start_date` varchar(20) NOT NULL,
  `ttl` varchar(20) NOT NULL,
  `request_id` varchar(20) NOT NULL,
  `request_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `active` TINYINT(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8;

-- Dumping data for table bte.amazon_report_request: ~72 rows (approximately)
/*!40000 ALTER TABLE `amazon_report_request` DISABLE KEYS */;
INSERT INTO `amazon_report_request` (`id`, `store`, `report_type`, `filename`, `start_date`, `ttl`, `request_id`) VALUES
	(1, 'CA', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__CA.txt', '', '2 hours', ''),
	(2, 'US', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__US.txt', '', '2 hours', ''),
	(3, 'CN', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__CN.txt', '', '2 hours', ''),
	(4, 'MX', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__MX.txt', '', '2 hours', ''),
	(5, 'UK', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__UK.txt', '', '2 hours', ''),
	(6, 'AU', '_GET_MERCHANT_LISTINGS_DATA_LITER_', 'MERCHANT_LISTINGS_DATA_LITER__AU.txt', '', '2 hours', ''),
	(7, 'CA', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__CA.txt', '', '2 hours', ''),
	(8, 'US', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__US.txt', '', '2 hours', ''),
	(9, 'CN', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__CN.txt', '', '2 hours', ''),
	(10, 'MX', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__MX.txt', '', '2 hours', ''),
	(11, 'UK', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__UK.txt', '', '2 hours', ''),
	(12, 'AU', '_GET_AFN_INVENTORY_DATA_', 'AFN_INVENTORY_DATA__AU.txt', '', '2 hours', ''),
	(13, 'CA', '_GET_ORDERS_DATA_', 'ORDERS_DATA__CA.txt', '', '2 hours', ''),
	(14, 'US', '_GET_ORDERS_DATA_', 'ORDERS_DATA__US.txt', '', '2 hours', ''),
	(15, 'CN', '_GET_ORDERS_DATA_', 'ORDERS_DATA__CN.txt', '', '2 hours', ''),
	(16, 'MX', '_GET_ORDERS_DATA_', 'ORDERS_DATA__MX.txt', '', '2 hours', ''),
	(17, 'UK', '_GET_ORDERS_DATA_', 'ORDERS_DATA__UK.txt', '', '2 hours', ''),
	(18, 'AU', '_GET_ORDERS_DATA_', 'ORDERS_DATA__AU.txt', '', '2 hours', ''),
	(19, 'CA', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__CA.txt', '', '2 hours', ''),
	(20, 'US', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__US.txt', '', '2 hours', ''),
	(21, 'CN', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__CN.txt', '', '2 hours', ''),
	(22, 'MX', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__MX.txt', '', '2 hours', ''),
	(23, 'UK', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__UK.txt', '', '2 hours', ''),
	(24, 'AU', '_GET_MERCHANT_LISTINGS_DATA_', 'MERCHANT_LISTINGS_DATA__AU.txt', '', '2 hours', ''),
	(25, 'CA', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__CA.txt', '-15 days', '10 minutes', ''),
	(26, 'US', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__US.txt', '-15 days', '10 minutes', ''),
	(27, 'CN', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__CN.txt', '-15 days', '10 minutes', ''),
	(28, 'MX', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__MX.txt', '-15 days', '10 minutes', ''),
	(29, 'UK', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__UK.txt', '-15 days', '10 minutes', ''),
	(30, 'AU', '_GET_FLAT_FILE_ORDERS_DATA_', 'FLAT_FILE_ORDERS_DATA__AU.txt', '-15 days', '10 minutes', ''),
	(31, 'CA', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__CA.txt', '-60 days', '2 hours', ''),
	(32, 'US', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__US.txt', '-60 days', '2 hours', ''),
	(33, 'CN', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__CN.txt', '-60 days', '2 hours', ''),
	(34, 'MX', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__MX.txt', '-60 days', '2 hours', ''),
	(35, 'UK', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__UK.txt', '-60 days', '2 hours', ''),
	(36, 'AU', '_GET_AMAZON_FULFILLED_SHIPMENTS_DATA_', 'AMAZON_FULFILLED_SHIPMENTS_DATA__AU.txt', '-60 days', '2 hours', ''),
	(37, 'CA', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__CA.txt', '-7 days', '2 hours', ''),
	(38, 'US', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__US.txt', '-7 days', '2 hours', ''),
	(39, 'CN', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__CN.txt', '-7 days', '2 hours', ''),
	(40, 'MX', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__MX.txt', '-7 days', '2 hours', ''),
	(41, 'UK', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__UK.txt', '-7 days', '2 hours', ''),
	(42, 'AU', '_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_', 'FLAT_FILE_ACTIONABLE_ORDER_DATA__AU.txt', '-7 days', '2 hours', ''),
	(43, 'CA', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__CA.txt', '-7 days', '2 hours', ''),
	(44, 'US', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__US.txt', '-7 days', '2 hours', ''),
	(45, 'CN', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__CN.txt', '-7 days', '2 hours', ''),
	(46, 'MX', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__MX.txt', '-7 days', '2 hours', ''),
	(47, 'UK', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__UK.txt', '-7 days', '2 hours', ''),
	(48, 'AU', '_GET_REFERRAL_FEE_PREVIEW_REPORT_', 'REFERRAL_FEE_PREVIEW_REPORT__AU.txt', '-7 days', '2 hours', ''),
	(49, 'CA', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__CA.txt', '-7 days', '2 hours', ''),
	(50, 'US', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__US.txt', '-7 days', '2 hours', ''),
	(51, 'CN', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__CN.txt', '-7 days', '2 hours', ''),
	(52, 'MX', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__MX.txt', '-7 days', '2 hours', ''),
	(53, 'UK', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__UK.txt', '-7 days', '2 hours', ''),
	(54, 'AU', '_GET_FLAT_FILE_PAYMENT_SETTLEMENT_DATA_', 'FLAT_FILE_PAYMENT_SETTLEMENT_DATA__AU.txt', '-7 days', '2 hours', ''),
	(55, 'CA', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__CA.txt', '-30 days', '2 hours', ''),
	(56, 'US', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__US.txt', '-30 days', '2 hours', ''),
	(57, 'CN', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__CN.txt', '-30 days', '2 hours', ''),
	(58, 'MX', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__MX.txt', '-30 days', '2 hours', ''),
	(59, 'UK', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__UK.txt', '-30 days', '2 hours', ''),
	(60, 'AU', '_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_', 'FBA_ESTIMATED_FBA_FEES_TXT_DATA__AU.txt', '-30 days', '2 hours', ''),
	(61, 'CA', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__CA.txt', '', '2 hours', ''),
	(62, 'US', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__US.txt', '', '2 hours', ''),
	(63, 'CN', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__CN.txt', '', '2 hours', ''),
	(64, 'MX', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__MX.txt', '', '2 hours', ''),
	(65, 'UK', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__UK.txt', '', '2 hours', ''),
	(66, 'AU', '_GET_MERCHANT_LISTINGS_ALL_DATA_', 'MERCHANT_LISTINGS_ALL_DATA__AU.txt', '', '2 hours', ''),
	(67, 'CA', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__CA.txt', '-7 days', '2 hours', ''),
	(68, 'US', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__US.txt', '-7 days', '2 hours', ''),
	(69, 'CN', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__CN.txt', '-7 days', '2 hours', ''),
	(70, 'MX', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__MX.txt', '-7 days', '2 hours', ''),
	(71, 'UK', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__UK.txt', '-7 days', '2 hours', ''),
	(72, 'AU', '_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_', 'FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA__AU.txt', '-7 days', '2 hours', '');
/*!40000 ALTER TABLE `amazon_report_request` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
