-- phpMyAdmin SQL Dump
-- version 4.7.6
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 20, 2018 at 01:08 AM
-- Server version: 10.1.30-MariaDB
-- PHP Version: 7.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

--
-- Database: `wedge`
--

-- --------------------------------------------------------

--
-- Table structure for table `accounts`
--

CREATE TABLE `accounts` (
  `ID` binary(16) NOT NULL,
  `ParentID` binary(16) NOT NULL,
  `Name` varchar(175) COLLATE utf8mb4_unicode_ci NOT NULL,
  `BookID` binary(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `aseets_values`
--

CREATE TABLE `aseets_values` (
  `ID` binary(16) NOT NULL,
  `AssetID` binary(16) NOT NULL,
  `BaseID` binary(16) NOT NULL,
  `AssetInBase` float NOT NULL,
  `LocalDate` int(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `assets`
--

CREATE TABLE `assets` (
  `ID` binary(16) NOT NULL,
  `Name` varchar(175) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Code` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'EX: BRL, BTC, USD, NASQ:GOOG, B3:PETR4',
  `BookID` binary(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `ID` binary(16) NOT NULL,
  `Name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Password` varbinary(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `ID` binary(16) NOT NULL,
  `AssetID` binary(16) NOT NULL,
  `TransactionID` binary(16) NOT NULL,
  `Name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Qty` float NOT NULL,
  `UnitCost` bigint(50) NOT NULL,
  `TotalCost` bigint(50) NOT NULL,
  `PeriodStart` int(8) NOT NULL COMMENT 'Y*1E4 + M*1E2 + D\nEx: 20171231 = 2017-12-31',
  `PeriodEnd` int(8) NOT NULL COMMENT 'Y*1E4 + M*1E2 + D\nEx: 20171231 = 2017-12-31'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `movements`
--

CREATE TABLE `movements` (
  `ID` binary(16) NOT NULL,
  `AccountID` binary(16) NOT NULL,
  `AssetID` binary(16) NOT NULL,
  `TransactionID` binary(16) NOT NULL,
  `Amount` bigint(50) NOT NULL,
  `Status` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'P' COMMENT '(P)Planned;(D)Done;(C)Cancelled',
  `LocalDate` int(8) NOT NULL COMMENT 'Y*1E4 + M*1E2 + D\nEx: 20171231 = 2017-12-31',
  `Notes` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Stand-in structure for view `movements_view`
-- (See below for the actual view)
--
CREATE TABLE `movements_view` (
`MovementID` binary(16)
,`TransactionID` binary(16)
,`AccountID` binary(16)
,`AccountParentID` binary(16)
,`AccountBookID` binary(16)
,`LocalDate` int(8)
,`Status` char(1)
,`AccountName` varchar(175)
,`AssetID` binary(16)
,`AssetCode` varchar(45)
,`Amount` bigint(50)
);

-- --------------------------------------------------------

--
-- Table structure for table `tags`
--

CREATE TABLE `tags` (
  `Id` int(11) NOT NULL,
  `Table` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ItemUUID` binary(16) NOT NULL,
  `Tag` varchar(160) COLLATE utf8mb4_unicode_ci NOT NULL,
  `BookUUID` binary(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `ID` binary(16) NOT NULL,
  `Name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `LocalDate` int(8) NOT NULL,
  `BookID` binary(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Structure for view `movements_view`
--
DROP TABLE IF EXISTS `movements_view`;

CREATE VIEW `movements_view`  AS  select `movements`.`ID` AS `MovementID`,`movements`.`TransactionID` AS `TransactionID`,`movements`.`AccountID` AS 
`AccountID`,`accounts`.`ParentID` AS `AccountParentID`,`accounts`.`BookID` AS `AccountBookID`,`movements`.`LocalDate` AS `LocalDate`,`movements`.`Status` AS 
`Status`,`accounts`.`Name` AS `AccountName`,`movements`.`AssetID` AS `AssetID`,`assets`.`Code` AS `AssetCode`,`movements`.`Amount` AS `Amount` from ((`movements` join 
`accounts` on((`movements`.`AccountID` = `accounts`.`ID`))) join `assets` on((`movements`.`AssetID` = `assets`.`ID`))) ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `accounts`
--
ALTER TABLE `accounts`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD UNIQUE KEY `Name_UNIQ` (`Name`,`BookID`),
  ADD KEY `fk_accounts_1_idx` (`ParentID`),
  ADD KEY `fk_accounts_2_idx` (`BookID`);

--
-- Indexes for table `aseets_values`
--
ALTER TABLE `aseets_values`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD UNIQUE KEY `UNIQ` (`BaseID`,`AssetID`,`LocalDate`),
  ADD KEY `fk_aseets_value_1_idx` (`AssetID`);

--
-- Indexes for table `assets`
--
ALTER TABLE `assets`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD UNIQUE KEY `Name_UNIQ` (`Code`,`BookID`) USING BTREE,
  ADD KEY `fk_assets_1_idx` (`BookID`);

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD KEY `fk_itens_1_idx` (`TransactionID`);

--
-- Indexes for table `movements`
--
ALTER TABLE `movements`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD KEY `fk_movements_1_idx` (`AssetID`),
  ADD KEY `fk_movements_2_idx` (`AccountID`),
  ADD KEY `fk_movements_3_idx` (`TransactionID`);

--
-- Indexes for table `tags`
--
ALTER TABLE `tags`
  ADD PRIMARY KEY (`Id`),
  ADD UNIQUE KEY `UNIQ` (`ItemUUID`,`Tag`,`BookUUID`),
  ADD KEY `fk_tags_1_idx` (`BookUUID`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UUID_UNIQUE` (`ID`),
  ADD KEY `Idx_BookID` (`BookID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tags`
--
ALTER TABLE `tags`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

