<?php require_once '_header.tpl'; ?>

	<h1><a target="_blank" href="http://google.com/search?q=<?=str_replace(' | ', ' ', $film->data->h1_title)?> смотреть фильм онлайн"><?=$film->data->h1_title?></a></h1>

	<?php if (isset($film->data->kinopoisk)) { ?><a href="<?=$film->data->kinopoisk?>" target="_blank">KINOPOISK</a><br /><?php } ?>

	<?php if (isset($film->data->imdb)) { ?><a href="<?=$film->data->imdb?>" target="_blank">IMDB</a><br /><?php } ?>

	<br />
	<button type="button" onclick="location.reload(); return false;">Get Film!</button>
	<br /><br />
	<img src="<?=$film->data->image_url?>" alt="<?=$film->data->h1_title?>" title="<?=$film->data->h1_title?>" />

<?php require_once '_footer.tpl'; ?>
