<?php

/*
*  Instance Configuration
*  ----------------------
*  Edit this file and not config.php for imageboard configuration.
*
*  You can copy values from config.php (defaults) and paste them here.
*/

	//$config['smart_build'] = true;
	$config['try_smarter'] = false;
	$config['spoiler_images'] = true;
	$config['field_disable_reply_subject'] = true;
	$config['strip_exif'] = true;
	$config['image_reject_repost'] = false;
//	$config['poster_ids'] = true;
	$config['thread_subject_in_title'] = true;

	$config['boards'] = array(
		array('fren', 'test')
	);
	$config['page_nav_top'] = true;
	$config['additional_javascript'][] = 'js/local-time.js';
	$config['additional_javascript'][] = 'js/auto-reload.js';
	$config['additional_javascript'][] = 'js/catalog-search.js'; // for catalog page
	$config['additional_javascript'][] = 'js/catalog.js';        // for catalog page
	$config['additional_javascript'][] = 'js/inline-expanding.js'; // for catalog page

	$config['db']['server'] = 'localhost';
	$config['db']['database'] = 'vichan';
	$config['db']['prefix'] = '';
	$config['db']['user'] = 'vichan-user';
	$config['db']['password'] = 'vichan-pass';

	$config['cookies']['mod'] = 'mod';
	$config['cookies']['salt'] = 'OSSL.QglincFaOBC+wMNnC3+T2VpGyBxAAYsTyM8kxfLacO3ZJIFT+W0wz419Z2B0VsUy9/8dYoA+0yJwYxrYHUN60lSx0x0S5jme7k3JbtJdaCMi/R1sJbfKtikM3rPpMGLbJZrMhycfZ9INeXCj5Lgr4FDoMmVdgQ9J7hKAoKd7Lwo=';

	$config['flood_time'] = 10;
	$config['flood_time_ip'] = 120;
	$config['flood_time_same'] = 30;
	$config['max_body'] = 1800;
	$config['reply_limit'] = 250;
	$config['max_links'] = 20;
	$config['max_filesize'] = 10485760;
	$config['thumb_width'] = 255;
	$config['thumb_height'] = 255;
	$config['max_width'] = 10000;
	$config['max_height'] = 10000;
	$config['threads_per_page'] = 10;
	$config['max_pages'] = 10;
	$config['threads_preview'] = 5;
	$config['root'] = '/';
	$config['secure_trip_salt'] = 'OSSL.zNnRzKyicgroCHnH7KH7S4u8YobUHJv4tyGC20yyScd7LEkib00orcvr6qi9Yv08ugn0Q0KV2JIo3E2ccrLBbup1vGFS/wEheuZ2xLFi6YR79jR5ZvAdk/v0VYXYTVDDKHBx6Tg8KUwqyMXhOy9N4xJ5POzgNTkqfv/HGJ2WULo=';

	$config['thumb_method'] = 'gm';
	$config['gnu_md5'] = '1';


// Changes made via web editor by "admin" @ Sat, 13 Nov 2021 13:53:18 -0800:
$config['global_message'] = '';

event_handler('post', function($post) {
	$post->name = trim($_SERVER['REMOTE_USER']);
	if (!isset($post->name)) {
		return "Username not provided in header";
	}

	if (preg_match("/^[a-zA-Z]+$/", $post->name)) { // usernames should match this, otherwise we risk injection
		$cmd = 'grep -A 1 "^' . $post->name . ':" /var/www/vichan/vichan-users | tail -n 1 | cut -d "#" -f2';

		$prettyname = trim(shell_exec($cmd));

		// return "Foo: " . $prettyname;
		if (isset($prettyname)) {
			$post->name = $prettyname;
		}
	}

});

// Changes made via web editor by "admin" @ Sat, 13 Nov 2021 14:05:35 -0800:
$config['verbose_errors'] = false;

// Changes made via web editor by "admin" @ Sat, 13 Nov 2021 14:42:16 -0800:
$config['force_image_op'] = false;

$config['field_disable_email'] = true;
$config['country_flags'] = true;
$config['allow_no_country'] = true;

