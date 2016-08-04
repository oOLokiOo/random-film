<?php

mb_internal_encoding('UTF-8');
error_reporting(E_ALL);
ini_set('error_reporting', E_ALL);


const ROOT = __DIR__;

require_once ROOT.'/../parser-kinopoisk/v3.0/inc/Parser/KinopoiskParser.php';
require_once ROOT.'/inc/RandomFilm/RandomFilm.php';

/* *** PROJECT USER XML STRUCTURE ***

<?xml version="1.0" encoding="UTF-8"?>
<movies>
	<film>
		<ru>test movie</ru>
		<en>тест</en>
		<year>1234</year>
		<kinopoisk>http://www.kinopoisk.ru/xxx</kinopoisk>
		<imdb>http://www.imdb.com/xxx</imdb>
	</film>
	<film>
		<ru>test movie 2</ru>
		<en>тест</en>
		<year>1234</year>
		<kinopoisk>http://www.kinopoisk.ru/xxx2</kinopoisk>
		<imdb>http://www.imdb.com/xxx2</imdb>
	</film>
</movies>
*/

$result 		= null;
$USER_XML_PATH 	= '../../../users/1/films.xml'; // TODO: link User Class here

$app = new RandomFilm($USER_XML_PATH);

require_once ROOT.'/tpl/index.tpl';
exit();
