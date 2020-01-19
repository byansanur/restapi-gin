-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 19, 2020 at 07:26 AM
-- Server version: 10.1.39-MariaDB
-- PHP Version: 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gogo`
--

-- --------------------------------------------------------

--
-- Table structure for table `notif`
--

CREATE TABLE `notif` (
  `id` int(11) NOT NULL,
  `id_sos_jamaah` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `notif`
--

INSERT INTO `notif` (`id`, `id_sos_jamaah`) VALUES
(1, 1),
(2, 3),
(3, 4),
(4, 5),
(5, 6),
(6, 7),
(7, 8),
(8, 9),
(9, 10),
(10, 11),
(11, 12),
(12, 13),
(13, 14),
(14, 15),
(15, 16),
(16, 17),
(17, 18),
(18, 19),
(19, 20);

-- --------------------------------------------------------

--
-- Table structure for table `rekomendasi`
--

CREATE TABLE `rekomendasi` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `foto` text NOT NULL,
  `lat` varchar(255) NOT NULL,
  `lng` varchar(255) NOT NULL,
  `rating` varchar(255) NOT NULL,
  `id_type` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `rekomendasi`
--

INSERT INTO `rekomendasi` (`id`, `nama`, `alamat`, `foto`, `lat`, `lng`, `rating`, `id_type`, `created_at`) VALUES
(1, 'Hilton Suites Makkah', 'Ibrahim Al Khalil, Jabal Omar? Mecca 21955, Arab Saudi / ??????? ??????? Jabal Omar? ??? 21955', '86f3eb914734c3b82b1e28e13315e4d7\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20200111234552.6rGGgzPVCq4RDxPQfunnamed(1).jpg', '21.421188', '39.821682', '4,6', 1, '2020-01-11 16:45:52'),
(2, 'Conrad Makkah', 'Jabal Omar, Ibrahim Al Khalil? Mecca 21955, Arab Saudi', '218a9e550901cc76cd53ddcb2d9722aa\\e6rYbxNN9SBmUutoXgTQcvMvjxcEEF5eHBqpLmum20200111234746.0tmNEjWYbP21Y4ZGtunnamed(2).jpg', '21.419410', '39.822407', '4,6', 1, '2020-01-11 16:47:46'),
(3, 'Makkah Hotel', 'Ibrahim Al Khalil, Al Hajlah, Mecca 24231, Arab Saudi / ??????? ??????? ??????? ??? 24231', 'a2b18650ff600d5a0ad2d92b846c581a\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20200111234226.3r9DviyJR9tAiNTzMunnamed.jpg', '21.420308', '39.823828', '4,4', 1, '2020-01-11 16:42:26'),
(4, 'Al Meraaj Restaurant Conrad Makkah', 'Ibrahim Al Khalil, Misfalah, ??, Mecca Arab Saudi / ??????? ??????? ???????? ??, ??? ', '093aa12721cec3577142e93048a98a40\\YUkQTt80Ce3SO7NGm0Qp64x0QMZCRgBFqTHbi5EY20200111235116.0tRW3x7liGOBWFFJtunnamed(3).jpg', '21.419726', '39.822300', '4,6', 2, '2020-01-11 16:51:16'),
(5, 'Grapari Makkah', '????? ??????? 4100 First Ring Rd, Ajyad, Mecca 24231, Arab Saudi', '15b52309e460af38edafafcee89cd95d\\FhVNaB0NydcF08ehQsSa46PcwTWf8UKrAGK0WC0F20200111235408.6l7YECzHHIfTyyp0munnamed(3).jpg', '21.418880', '39.824765', '4,2', 2, '2020-01-11 16:54:08'),
(6, 'Zamzam', '4256 King Abdul Aziz Rd, Mecca 24231, Arab Saudi', 'da72860ec742f787991bb14521bcbf77\\7riqSOlpahM5GmF4Qg5J4oCGG3HgODjcEZ6apSWO20200111235738.6mrRKxNjJFR5Q4eXkunnamed(5).jpg', '21.422659', '39.826522', '4,8', 3, '2020-01-11 16:57:38');

-- --------------------------------------------------------

--
-- Table structure for table `sos_jamaah`
--

CREATE TABLE `sos_jamaah` (
  `id` int(11) NOT NULL,
  `message` text NOT NULL,
  `id_users_sender` int(11) NOT NULL,
  `lat` text NOT NULL,
  `lng` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `sos_jamaah`
--

INSERT INTO `sos_jamaah` (`id`, `message`, `id_users_sender`, `lat`, `lng`, `created_at`) VALUES
(1, '', 12, '65363455', '6363463', '2020-01-05 08:22:52'),
(2, '', 11, '642365534', '65478377648', '2020-01-05 09:22:57'),
(3, '', 11, '642365534', '65478377648', '2020-01-05 09:33:04'),
(4, '', 11, '6563466', '654783736367648', '2020-01-05 09:42:08'),
(5, '', 11, '635345345545', '66354', '2020-01-05 13:03:06'),
(6, 'test sos', 12, '21.422659', '39.826522', '2020-01-12 07:34:45'),
(7, 'test sos again', 12, '21.422659', '39.826522', '2020-01-12 07:35:40'),
(8, 'test sos again again', 12, '21.422659', '39.826522', '2020-01-12 11:04:42'),
(9, 'help', 11, '0.0', '0.0', '2020-01-12 11:35:04'),
(10, 'jjj', 11, '0.0', '0.0', '2020-01-12 11:36:22'),
(11, 'hffjk', 11, '0.0', '0.0', '2020-01-12 11:50:41'),
(12, 'hhg', 12, '-6.1424697', '106.8869963', '2020-01-16 08:18:58'),
(13, 'gffd', 12, '-6.1424697', '106.8869963', '2020-01-16 08:20:52'),
(14, 'bb', 12, '-6.1424697', '106.8869963', '2020-01-16 08:22:03'),
(15, 'tezy', 12, '-6.2550543', '107.0877185', '2020-01-16 16:29:22'),
(16, 'fdfgh', 12, '-6.1424749', '106.8869676', '2020-01-17 07:34:25'),
(17, 'yihfx', 11, '-6.2595056', '107.0883411', '2020-01-17 17:29:21'),
(18, 'tolong saya tersesat', 11, '-6.2594927', '107.0883323', '2020-01-17 17:59:08'),
(19, 'tolong bapak saya kan suka lupa wkwkw', 12, '-6.2595054', '107.0883412', '2020-01-18 02:16:25'),
(20, 'help', 11, '-6.2595004', '107.0883402', '2020-01-19 06:04:46');

-- --------------------------------------------------------

--
-- Table structure for table `tb_privileges`
--

CREATE TABLE `tb_privileges` (
  `id` int(11) NOT NULL,
  `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `tb_privileges`
--

INSERT INTO `tb_privileges` (`id`, `role`) VALUES
(1, 'admin'),
(2, 'petugas'),
(3, 'jamaah');

-- --------------------------------------------------------

--
-- Table structure for table `typerekom`
--

CREATE TABLE `typerekom` (
  `id` int(11) NOT NULL,
  `type_rekom` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `typerekom`
--

INSERT INTO `typerekom` (`id`, `type_rekom`) VALUES
(1, 'Hotel & Ressort'),
(2, 'Halal Food'),
(3, 'Travel');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tgl_lahir` varchar(15) NOT NULL,
  `no_ktp` varchar(17) NOT NULL,
  `no_hp` varchar(17) NOT NULL,
  `no_visa` varchar(255) NOT NULL,
  `no_passpor` varchar(255) NOT NULL,
  `foto` varchar(255) NOT NULL,
  `id_privileges` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `lat` text NOT NULL,
  `lng` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `username`, `password`, `tgl_lahir`, `no_ktp`, `no_hp`, `no_visa`, `no_passpor`, `foto`, `id_privileges`, `created_at`, `lat`, `lng`) VALUES
(1, 'ica tester', 'test', '8XaAAOiU-OYnftQNMzEvPAgD1ghwNO9PYzpwfC8fyV8', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(2, 'ica1', 'ica1', '', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(3, 'ica12', 'ica12', '', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(4, 'ica123', 'ica123', '', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(5, 'ica1234', 'ica1234', '4aABSCgO6hPUrhTnUco4H9iBoV2yy-Brl_oDoB3hfgc', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(6, 'ica tester', 'test', 'MIr_eHkanblKmB6OnZ3nIIh4itV-Ibe_DpIXAgEaNo4', '0000-00-', '0', '0', '', '', '', 0, '2019-12-06 14:58:17', '', ''),
(7, 'data lengkap', 'data', 'uP8XbdGZWumwFmnZRwQV88jtQsT5FGq-ieJ4la_W2yc', '1999-01-', '2147483647', '887878529', '23957203572727298', '73647hhfuieyrui73', '8d777f385d3dfec8815d20f7496026dc\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191206224047.0r9DviyJR9tAiNTzMselamatpagi.jpg', 3, '2020-01-17 17:55:07', '-6.2594927', '107.0883323'),
(9, 'admin', 'admin', 'ijqz7PEJDUT3o0ZILnaWAutvGDvT49oA68DyLXIpsOY', '1993-04-', '2147483647', '2147483647', '323732757584759748', 'er784577wgfw9', '21232f297a57a5a743894a0e4a801fc3\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191206231140.0r9DviyJR9tAiNTzMselamatpagi.jpg', 1, '2019-12-06 16:11:40', '', ''),
(10, 'petugas', 'petugas', 'pqNqVbY_0Sa_9roso6wDABNEEzRzOs-BFD6lEu0hmts', '1993-04-', '2147483647', '2147483647', '4273657263924572387', '27678rebr', 'afb91ef692fd08c445e8cb1bab2ccf9c\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20191206231248.0rGGgzPVCq4RDxPQfselamatpagi.jpg', 2, '2019-12-06 16:12:48', '', ''),
(11, 'Husni alhayubi bagas', 'husni', 'SZUNNnfMM4Zx9fT24ksvtWpt_i6b9wDXct7B_lWaZrU', '1992-08-4', '3216070908961118', '81264782800', '3216070908961118', '124638fg5', '143196712ca8d8714a875522c5957a6d\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191215012326.6r9DviyJR9tAiNTzMIMG_20191215_010155.jpg', 3, '2020-01-19 06:15:59', '-6.2594959', '107.0883361'),
(12, 'Yahya ahmad bunaryah', 'yahya', 'CVH4j9IyT_KX2VYQycnqr8psHmH9p4tPDy3DtuXmnig', '1972-08-4', '368374628274859', '81573882800', '368374628274859', '6589346fg5', '59202463fd4c312b063293b88f6063b2\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191215042308.7r9DviyJR9tAiNTzMbackground.png', 3, '2020-01-18 02:16:25', '-6.2595054', '107.0883412'),
(13, 'Beni aljamaludin kamal', 'beni', 'jFEg3sD-xpSAUJG6wrnUyJ2a-VwFumzkdWANB3y9oZw', '22-12-1996', '3217876472639469', '81243739302', '3217876472639469', '62865894', 'b94ce3c426a5ab6032624ab62a2b0b95\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20200104162337.6r9DviyJR9tAiNTzMgaya.png', 2, '2020-01-04 09:23:37', '', ''),
(14, 'Ratbyansa nur', 'byan', 'QZRbVrXwvRcfMRZKfoSKvK9EjQui1yzypUzcIxA9HyE', '22-12-1996', '3216070908960004', '81243739302', '3216070908960004', 'ina563805888', 'c65ff08f4f742ccd13e0a1180254d8db\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20200104162707.2rGGgzPVCq4RDxPQfgaya.png', 2, '2020-01-04 09:27:07', '', ''),
(16, 'yoga saepulloh', 'yoga', 'SHb0eyqd7EsjLgCEmKrOqaAyZraA99vG22UORyYPvlU', '04-10-1999', '3216070958030005', '81264567089', '3216070958030005', 'gfjkt12r545465555', '', 2, '2020-01-19 06:16:10', '-6.2594959', '107.0883361');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `notif`
--
ALTER TABLE `notif`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `rekomendasi`
--
ALTER TABLE `rekomendasi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sos_jamaah`
--
ALTER TABLE `sos_jamaah`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tb_privileges`
--
ALTER TABLE `tb_privileges`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `typerekom`
--
ALTER TABLE `typerekom`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `notif`
--
ALTER TABLE `notif`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- AUTO_INCREMENT for table `rekomendasi`
--
ALTER TABLE `rekomendasi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `sos_jamaah`
--
ALTER TABLE `sos_jamaah`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `tb_privileges`
--
ALTER TABLE `tb_privileges`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `typerekom`
--
ALTER TABLE `typerekom`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
