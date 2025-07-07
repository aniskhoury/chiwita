/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19  Distrib 10.11.13-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: chiwita
-- ------------------------------------------------------
-- Server version	10.11.13-MariaDB-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `administracion`
--

DROP TABLE IF EXISTS `administracion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `administracion` (
  `idadministracion` bigint(20) NOT NULL AUTO_INCREMENT,
  `idusuario` bigint(20) NOT NULL,
  `cargo` varchar(128) NOT NULL,
  `nivel` bigint(20) NOT NULL,
  PRIMARY KEY (`idadministracion`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `administracion`
--

LOCK TABLES `administracion` WRITE;
/*!40000 ALTER TABLE `administracion` DISABLE KEYS */;
INSERT INTO `administracion` VALUES
(1,1,'Creador',0);
/*!40000 ALTER TABLE `administracion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `canales`
--

DROP TABLE IF EXISTS `canales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `canales` (
  `idcanal` bigint(20) NOT NULL AUTO_INCREMENT,
  `idcategoria` bigint(20) NOT NULL,
  `idservidor` bigint(20) NOT NULL,
  `nombrecanal` varchar(32) NOT NULL,
  PRIMARY KEY (`idcanal`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `canales`
--

LOCK TABLES `canales` WRITE;
/*!40000 ALTER TABLE `canales` DISABLE KEYS */;
INSERT INTO `canales` VALUES
(1,1,1,'Introducción');
/*!40000 ALTER TABLE `canales` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categoriacanales`
--

DROP TABLE IF EXISTS `categoriacanales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `categoriacanales` (
  `idcategoria` bigint(20) NOT NULL AUTO_INCREMENT,
  `idservidor` bigint(20) NOT NULL,
  `nombreCategoria` varchar(32) NOT NULL,
  PRIMARY KEY (`idcategoria`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categoriacanales`
--

LOCK TABLES `categoriacanales` WRITE;
/*!40000 ALTER TABLE `categoriacanales` DISABLE KEYS */;
INSERT INTO `categoriacanales` VALUES
(1,1,'Juegos'),
(2,1,'Charlar'),
(3,1,'General');
/*!40000 ALTER TABLE `categoriacanales` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mensajes`
--

DROP TABLE IF EXISTS `mensajes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mensajes` (
  `idmensaje` bigint(20) NOT NULL AUTO_INCREMENT,
  `idusuario` bigint(20) NOT NULL,
  `idservidor` bigint(20) NOT NULL,
  `canal` varchar(128) NOT NULL,
  `mensaje` varchar(4096) NOT NULL,
  `fecha_mensaje` datetime NOT NULL,
  PRIMARY KEY (`idmensaje`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mensajes`
--

LOCK TABLES `mensajes` WRITE;
/*!40000 ALTER TABLE `mensajes` DISABLE KEYS */;
/*!40000 ALTER TABLE `mensajes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mensajesprivados`
--

DROP TABLE IF EXISTS `mensajesprivados`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mensajesprivados` (
  `idmensajesprivados` bigint(20) NOT NULL AUTO_INCREMENT,
  `idusuarioemisor` bigint(20) NOT NULL,
  `idusuarioreceptor` bigint(20) NOT NULL,
  `idservidor` bigint(20) NOT NULL,
  `mensaje` varchar(4096) NOT NULL,
  `fecha_mensajeprivado` datetime NOT NULL,
  PRIMARY KEY (`idmensajesprivados`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mensajesprivados`
--

LOCK TABLES `mensajesprivados` WRITE;
/*!40000 ALTER TABLE `mensajesprivados` DISABLE KEYS */;
/*!40000 ALTER TABLE `mensajesprivados` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `servidores`
--

DROP TABLE IF EXISTS `servidores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `servidores` (
  `idservidor` bigint(20) NOT NULL AUTO_INCREMENT,
  `nombreServidor` varchar(32) NOT NULL,
  `ipServidor` varchar(32) NOT NULL,
  PRIMARY KEY (`idservidor`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `servidores`
--

LOCK TABLES `servidores` WRITE;
/*!40000 ALTER TABLE `servidores` DISABLE KEYS */;
INSERT INTO `servidores` VALUES
(1,'Neptuno','localhost');
/*!40000 ALTER TABLE `servidores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `idusuario` bigint(20) NOT NULL AUTO_INCREMENT,
  `nick` varchar(16) NOT NULL,
  `contrasena` char(128) NOT NULL,
  `email` varchar(64) NOT NULL,
  `fecha_registro` datetime NOT NULL,
  `activado` tinyint(1) NOT NULL,
  `bot` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`idusuario`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
INSERT INTO `usuarios` VALUES
(1,'Anis','d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db','akhouryribas@gmail.com','2025-06-24 18:52:02',1,0);
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-07  6:11:44
