-- phpMyAdmin SQL Dump
-- version 3.5.8.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jun 20, 2015 at 10:38 AM
-- Server version: 5.5.42-37.1-log
-- PHP Version: 5.4.23

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `longsinc_brewdatabase`
--

-- --------------------------------------------------------

--
-- Table structure for table `Equipment`
--

DROP TABLE IF EXISTS `Equipment`;
CREATE TABLE IF NOT EXISTS `Equipment` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` int(11) NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `BatchSize` decimal(11,6) NOT NULL,
  `TunVolume` decimal(11,6) NOT NULL,
  `TunWeight` decimal(11,6) NOT NULL,
  `TunSpecificHeat` decimal(11,6) NOT NULL,
  `TopUpWater` decimal(11,6) NOT NULL,
  `TrubChillerLoss` decimal(11,6) NOT NULL,
  `EvapRate` decimal(11,6) NOT NULL,
  `Sulfate` decimal(11,6) NOT NULL,
  `CalcBoilVolume` tinyint(1) NOT NULL,
  `LauterDeadspace` decimal(11,6) NOT NULL,
  `TopUpKettle` decimal(11,6) NOT NULL,
  `HopUtilization` decimal(11,6) NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayBoilSize` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayBatchSize` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTunVolume` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTunWeight` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTopUpWater` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTrubChillerLoss` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayLauterDeadspace` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTopUpKettle` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `EquipmentLookup`
--

DROP TABLE IF EXISTS `EquipmentLookup`;
CREATE TABLE IF NOT EXISTS `EquipmentLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `EquipmentID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Fermentable`
--

DROP TABLE IF EXISTS `Fermentable`;
CREATE TABLE IF NOT EXISTS `Fermentable` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Amount` decimal(11,6) NOT NULL,
  `Yield` decimal(11,6) NOT NULL,
  `Color` decimal(11,6) NOT NULL,
  `AddAfterBoil` tinyint(1) NOT NULL,
  `Origin` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Supplier` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `CoarseFineDiff` decimal(11,6) NOT NULL,
  `Moisture` decimal(11,6) NOT NULL,
  `DiastaticPower` decimal(11,6) NOT NULL,
  `Protein` decimal(11,6) NOT NULL,
  `MaxInBatch` decimal(11,6) NOT NULL,
  `ReccomendMash` tinyint(1) NOT NULL,
  `IBUGalPerLB` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayAmount` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Potential` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Inventory` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayColor` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `FermentableLookup`
--

DROP TABLE IF EXISTS `FermentableLookup`;
CREATE TABLE IF NOT EXISTS `FermentableLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `FermentableID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Hop`
--

DROP TABLE IF EXISTS `Hop`;
CREATE TABLE IF NOT EXISTS `Hop` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParenID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Origin` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Alpha` decimal(11,6) NOT NULL,
  `Amount` decimal(11,6) NOT NULL,
  `Use` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Time` decimal(11,6) NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Type` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Form` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Beta` decimal(11,6) NOT NULL,
  `HSI` decimal(11,6) NOT NULL,
  `Substitues` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Humulene` decimal(11,6) NOT NULL,
  `Caryophyllene` decimal(11,6) NOT NULL,
  `Cohumulone` decimal(11,6) NOT NULL,
  `Myrcene` decimal(11,6) NOT NULL,
  `DisplayAmount` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Inventory` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTime` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `HopLookup`
--

DROP TABLE IF EXISTS `HopLookup`;
CREATE TABLE IF NOT EXISTS `HopLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `HopID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `MashProfile`
--

DROP TABLE IF EXISTS `MashProfile`;
CREATE TABLE IF NOT EXISTS `MashProfile` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `GrainTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Notes` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `TunTemp` decimal(11,6) NOT NULL,
  `SpargeTemp` decimal(11,6) NOT NULL,
  `PH` decimal(11,6) NOT NULL,
  `TunWeight` decimal(11,6) NOT NULL,
  `TunSpecificHeat` decimal(11,6) NOT NULL,
  `EquipAdjust` tinyint(1) NOT NULL,
  `DisplayGrainTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTunTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplaySpargeTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTunWeight` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `MashProfile`
--

INSERT INTO `MashProfile` (`ID`, `ParentID`, `Name`, `Version`, `GrainTemp`, `Notes`, `TunTemp`, `SpargeTemp`, `PH`, `TunWeight`, `TunSpecificHeat`, `EquipAdjust`, `DisplayGrainTemp`, `DisplayTunTemp`, `DisplaySpargeTemp`, `DisplayTunWeight`) VALUES
(1, -1, 'Profile1', 1, '', '', '0.000000', '0.000000', '0.000000', '0.000000', '0.000000', 0, '', '', '', ''),
(2, -1, 'Profile2', 1, '', '', '0.000000', '0.000000', '0.000000', '0.000000', '0.000000', 0, '', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `MashProfileLookup`
--

DROP TABLE IF EXISTS `MashProfileLookup`;
CREATE TABLE IF NOT EXISTS `MashProfileLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `MashProfileID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `MashStep`
--

DROP TABLE IF EXISTS `MashStep`;
CREATE TABLE IF NOT EXISTS `MashStep` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Type` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `InfuseAmount` decimal(11,6) NOT NULL,
  `StepTemp` decimal(11,6) NOT NULL,
  `StepTime` decimal(11,6) NOT NULL,
  `RampTime` decimal(11,6) NOT NULL,
  `EndTemp` decimal(11,6) NOT NULL,
  `Description` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `WaterGrainRatio` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DecoctionAmt` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `InfuseTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayStepTemp` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayInfuseAmt` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `MashStep`
--

INSERT INTO `MashStep` (`ID`, `ParentID`, `Name`, `Version`, `Type`, `InfuseAmount`, `StepTemp`, `StepTime`, `RampTime`, `EndTemp`, `Description`, `WaterGrainRatio`, `DecoctionAmt`, `InfuseTemp`, `DisplayStepTemp`, `DisplayInfuseAmt`) VALUES
(1, -1, 'Step1', 1, '', '0.000000', '0.000000', '0.000000', '0.000000', '0.000000', '', '', '', '', '', ''),
(2, -1, 'Step2', 1, '', '0.000000', '0.000000', '0.000000', '0.000000', '0.000000', '', '', '', '', '', '');

-- --------------------------------------------------------

--
-- Table structure for table `MashStepLookup`
--

DROP TABLE IF EXISTS `MashStepLookup`;
CREATE TABLE IF NOT EXISTS `MashStepLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `MashProfileID` int(11) NOT NULL,
  `MashStepId` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `MashStepLookup`
--

INSERT INTO `MashStepLookup` (`ID`, `MashProfileID`, `MashStepId`) VALUES
(1, 1, 1),
(2, 1, 2);

-- --------------------------------------------------------

--
-- Table structure for table `Misc`
--

DROP TABLE IF EXISTS `Misc`;
CREATE TABLE IF NOT EXISTS `Misc` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Type` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Use` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Time` decimal(11,6) NOT NULL,
  `Amount` decimal(11,6) NOT NULL,
  `AmountIsWeight` tinyint(1) NOT NULL,
  `UseFor` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Notes` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayAmount` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Inventory` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTime` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `MiscLookup`
--

DROP TABLE IF EXISTS `MiscLookup`;
CREATE TABLE IF NOT EXISTS `MiscLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `MiscID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Recipe`
--

DROP TABLE IF EXISTS `Recipe`;
CREATE TABLE IF NOT EXISTS `Recipe` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Type` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Brewer` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `AsstBrewer` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `BatchSize` decimal(11,6) NOT NULL,
  `BoilSize` int(11) NOT NULL,
  `BoilTime` decimal(11,6) NOT NULL,
  `Efficiency` decimal(11,6) NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `TasteNotes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `TasteRating` decimal(11,6) NOT NULL,
  `OG` decimal(11,6) NOT NULL,
  `FG` decimal(11,6) NOT NULL,
  `FermentationStages` int(11) NOT NULL DEFAULT '1',
  `PrimaryAge` decimal(11,6) NOT NULL,
  `PrimaryTemp` decimal(11,6) NOT NULL,
  `SecondaryAge` decimal(11,6) NOT NULL,
  `SecondaryTemp` decimal(11,6) NOT NULL,
  `TertiaryAge` decimal(11,6) NOT NULL,
  `TertiaryTemp` decimal(11,6) NOT NULL,
  `Age` decimal(11,6) NOT NULL,
  `AgeTemp` decimal(11,6) NOT NULL,
  `Date` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Carbonation` decimal(11,6) NOT NULL,
  `ForceCarbonation` tinyint(1) NOT NULL,
  `PrimingSugarName` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `CarbonationTemp` decimal(11,6) NOT NULL,
  `PrimingSugarEquiv` decimal(11,6) NOT NULL,
  `KegPrimingFactor` decimal(11,6) NOT NULL,
  `EstOG` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `EstFG` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `EstColor` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `IBU` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `IBUMethod` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `EstABV` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `ABV` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `ActualEfficiency` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Calories` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayBatchSize` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayBoilSize` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayOG` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayFG` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayPrimaryTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplaySecondaryTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayTertiaryTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayAgeTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `CarbonationUsed` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayCarbTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Style`
--

DROP TABLE IF EXISTS `Style`;
CREATE TABLE IF NOT EXISTS `Style` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Category` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `CategoryNumber` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `StyleLetter` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `StyleGuide` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Type` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `OGMin` decimal(11,6) NOT NULL,
  `OGMax` decimal(11,6) NOT NULL,
  `FGMin` decimal(11,6) NOT NULL,
  `FGMax` decimal(11,6) NOT NULL,
  `IBUMin` decimal(11,6) NOT NULL,
  `IBUMax` decimal(11,6) NOT NULL,
  `ColorMin` decimal(11,6) NOT NULL,
  `ColorMax` decimal(11,6) NOT NULL,
  `CarbMin` decimal(11,6) NOT NULL,
  `CarbMax` decimal(11,6) NOT NULL,
  `ABVMin` decimal(11,6) NOT NULL,
  `ABVMax` decimal(11,6) NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Profile` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Ingredients` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Examples` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayOGMin` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayOGMax` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayFGMin` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayFGMax` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayColorMin` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DisplayColorMax` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `OGRange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `FGRange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `IBURange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `CarbRange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `ColorRange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `ABVRange` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `StyleLookup`
--

DROP TABLE IF EXISTS `StyleLookup`;
CREATE TABLE IF NOT EXISTS `StyleLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `StyleID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Water`
--

DROP TABLE IF EXISTS `Water`;
CREATE TABLE IF NOT EXISTS `Water` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Amount` decimal(11,6) NOT NULL,
  `Calcium` decimal(11,6) NOT NULL,
  `Bicarbonate` decimal(11,6) NOT NULL,
  `Sulfate` decimal(11,6) NOT NULL,
  `Chloride` decimal(11,6) NOT NULL,
  `Sodium` decimal(11,6) NOT NULL,
  `Magnesium` decimal(11,6) NOT NULL,
  `PH` decimal(11,6) NOT NULL,
  `Notes` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `DisplayAmount` varchar(256) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `WaterLookup`
--

DROP TABLE IF EXISTS `WaterLookup`;
CREATE TABLE IF NOT EXISTS `WaterLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `WaterID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `Yeast`
--

DROP TABLE IF EXISTS `Yeast`;
CREATE TABLE IF NOT EXISTS `Yeast` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ParentID` int(11) NOT NULL DEFAULT '-1',
  `Name` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Version` int(11) NOT NULL DEFAULT '1',
  `Type` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Form` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Amount` decimal(11,6) NOT NULL,
  `AmountIsWeight` tinyint(1) NOT NULL,
  `Laboratory` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `ProductID` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `MinTemperature` decimal(11,6) NOT NULL,
  `MaxTemperature` decimal(11,6) NOT NULL,
  `Flocculation` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Attenuation` decimal(11,6) NOT NULL,
  `Notes` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `BestFor` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `TimesCultured` int(11) NOT NULL DEFAULT '0',
  `MaxReuse` int(11) NOT NULL DEFAULT '1',
  `AddToSecondary` tinyint(1) NOT NULL,
  `DisplayAmount` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DispMinTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `DispMaxTemp` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Inventory` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `CultureDate` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `YeastLookup`
--

DROP TABLE IF EXISTS `YeastLookup`;
CREATE TABLE IF NOT EXISTS `YeastLookup` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RecipeID` int(11) NOT NULL,
  `YeastID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
