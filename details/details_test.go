package details

import (
	"testing"

	"github.com/gar-r/ngore/parse"
	"github.com/stretchr/testify/assert"
)

func TestParseDetails(t *testing.T) {

	t.Run("movie details", func(t *testing.T) {
		doc := parse.MustParse(t, movieHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "film", details.Type)
		assert.Equal(t, "Star Trek 9. - Űrlázadás (Star Trek: Insurrection)", details.Title)
		assert.Equal(t, "1998", details.ReleaseYear)
		assert.Equal(t, "Jonathan Frakes", details.Director)
		assert.Equal(t, "Patrick Stewart, Jonathan Frakes, Brent Spiner, LeVar Burton, Michael Dorn", details.Actors)
		assert.Equal(t, "USA", details.Country)
		assert.Equal(t, "akció kaland sci-fi thriller", details.Labels)
		assert.Equal(t, "6.3", details.ImdbRating)
		assert.Equal(t, "https://imdb.com/title/tt0120844", details.ImdbLink)
		assert.Equal(t, "103 perc", details.Length)
		assert.Equal(t, "http://www.port.hu/star_trek_9._-_urlazadas_star_trek_-_insurrection/pls/fi/films.film_page?i_film_id=3626", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/RPev3b57i0mfKBgw", details.CoverImage)
		assert.Equal(t, []string{
			"https://nc-img.burst.sh/EvjeNybRGtEi-BNx",
			"https://nc-img.burst.sh/6RGBnnrb7UVCOBwJ",
			"https://nc-img.burst.sh/xqlB6pY9VIKfRXdE",
		}, details.OtherImages)
	})

	t.Run("series details", func(t *testing.T) {
		doc := parse.MustParse(t, seriesHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "sorozat", details.Type)
		assert.Equal(t, "Star Trek: Voyager", details.Title)
		assert.Equal(t, "1995", details.ReleaseYear)
		assert.Equal(t, "David Livingston, Winrich Kolbe", details.Director)
		assert.Equal(t, "Kate Mulgrew, Robert Beltran, Roxann Dawson, Robert Duncan McNeill, Ethan\n\t\t\t\t\t\t\t\t\t\tPhillips", details.Actors)
		assert.Equal(t, "USA", details.Country)
		assert.Equal(t, "akció kaland sci-fi", details.Labels)
		assert.Equal(t, "7.7", details.ImdbRating)
		assert.Equal(t, "https://imdb.com/title/tt0112178", details.ImdbLink)
		assert.Equal(t, "45 perc", details.Length)
		assert.Equal(t, "http://www.port.hu/star_trek:_voyager/pls/w/films.film_page?i_film_id=26841", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/lLBo6Kb7hPpSwXvW?24340072", details.CoverImage)
		assert.Equal(t, []string{
			"https://nc-img.burst.sh/QoqB738Y0UzCd4kx?24340072",
			"https://nc-img.burst.sh/mDQ42-vPqU6fo4yW?24340072",
			"https://nc-img.burst.sh/2lwBDJ70pc7hr4N8?24340072",
		}, details.OtherImages)
	})

	t.Run("game details", func(t *testing.T) {
		doc := parse.MustParse(t, gameHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "játék", details.Type)
		assert.Equal(t, "Baldur's Gate - Enhanced Edition v2.6.5.0.45309 GOG", details.Title)
		assert.Equal(t, "", details.ReleaseYear)
		assert.Equal(t, "", details.Director)
		assert.Equal(t, "", details.Actors)
		assert.Equal(t, "", details.Country)
		assert.Equal(t, "", details.Labels)
		assert.Equal(t, "", details.ImdbRating)
		assert.Equal(t, "", details.ImdbLink)
		assert.Equal(t, "", details.Length)
		assert.Equal(t, "", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/jg_mpnRviRwfle8z?26990706", details.CoverImage)
		assert.Equal(t, []string{
			"https://nc-img.burst.sh/Jl3X9v9W6fyCD_7R?26990706",
			"https://nc-img.burst.sh/y7D4rJ-oPTbf74vV?26990706",
			"https://nc-img.burst.sh/6ZV_31yAqiJhwepM?26990706",
		}, details.OtherImages)
	})

	t.Run("music details", func(t *testing.T) {
		doc := parse.MustParse(t, musicHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "zene", details.Type)
		assert.Equal(t, "Linkin Park - Greatest Hits", details.Title)
		assert.Equal(t, "", details.ReleaseYear)
		assert.Equal(t, "", details.Director)
		assert.Equal(t, "", details.Actors)
		assert.Equal(t, "", details.Country)
		assert.Equal(t, "alternative metal", details.Labels)
		assert.Equal(t, "", details.ImdbRating)
		assert.Equal(t, "", details.ImdbLink)
		assert.Equal(t, "", details.Length)
		assert.Equal(t, "", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/DQ_2ZqxWT69io_yW", details.CoverImage)
		assert.Empty(t, details.OtherImages)
	})

	t.Run("app details", func(t *testing.T) {
		doc := parse.MustParse(t, appHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "program", details.Type)
		assert.Equal(t, "Total Commander v10.0 HUN-ENG x86-x64", details.Title)
		assert.Equal(t, "", details.ReleaseYear)
		assert.Equal(t, "", details.Director)
		assert.Equal(t, "", details.Actors)
		assert.Equal(t, "", details.Country)
		assert.Equal(t, "", details.Labels)
		assert.Equal(t, "", details.ImdbRating)
		assert.Equal(t, "", details.ImdbLink)
		assert.Equal(t, "", details.Length)
		assert.Equal(t, "", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/PxBzREgbSMlT-BbM?27057497", details.CoverImage)
		assert.Equal(t, []string{
			"https://nc-img.burst.sh/jvw_ZoJZkSrCzBNl?27057497",
			"https://nc-img.burst.sh/N-jB-ONG7CRfz_0v?27057497",
			"https://nc-img.burst.sh/AL9XkMlLKCLhlXZl?27057497",
		}, details.OtherImages)
	})

	t.Run("book details", func(t *testing.T) {
		doc := parse.MustParse(t, bookHtml)

		details := ParseDetails(doc)

		assert.Equal(t, "ebook", details.Type)
		assert.Equal(t, "George R. R. Martin - Királyok csatája - A Tűz és Jég dala II.", details.Title)
		assert.Equal(t, "", details.ReleaseYear)
		assert.Equal(t, "", details.Director)
		assert.Equal(t, "", details.Actors)
		assert.Equal(t, "", details.Country)
		assert.Equal(t, "fantasy", details.Labels)
		assert.Equal(t, "", details.ImdbRating)
		assert.Equal(t, "", details.ImdbLink)
		assert.Equal(t, "", details.Length)
		assert.Equal(t, "", details.OtherLink)
		assert.Equal(t, "https://nc-img.burst.sh/covers/EdeYZgjvCDrCLX-x", details.CoverImage)
		assert.Equal(t, []string{"https://nc-img.burst.sh/N-jX-93vQTPizB0v"}, details.OtherImages)
	})

}

const movieHtml = `
<div class="fobox_tartalom">
	<div class="torrent_reszletek_cim">Star Trek 9. - Űrlázadás</div>
	<div class="torrent_reszletek_konyvjelzo">
		<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);" onclick="konyvjelzo('1365642');"
			title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
		<a style="font-weight:normal;"
			href="torrents.php?action=addnews&amp;id=1365642&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
			title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
		<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=1365642">[bejelentés]</a>
	</div>
	<center>
		<table cellspacing="5" border="0">
			<tbody>
				<tr>
					<td> <a class="fancy_groups" rel="g1365642"
							href="https://nc-img.burst.sh/EvjeNybRGtEi-BNx"><img
								src="https://nc-img.burst.sh/6RGBnnrb7UmiwBwJ" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g1365642"
							href="https://nc-img.burst.sh/6RGBnnrb7UVCOBwJ"><img
								src="https://nc-img.burst.sh/xqlB6pY9VImCbXdE" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g1365642"
							href="https://nc-img.burst.sh/xqlB6pY9VIKfRXdE"><img
								src="https://nc-img.burst.sh/qjA_Rk7yVC9fl_-7" border="0"></a></td>
				</tr>
				<tr>
					<td align="center"><em>(720x304)</em></td>
					<td align="center"><em>(720x304)</em></td>
					<td align="center"><em>(720x304)</em></td>
				</tr>
			</tbody>
		</table>
	</center><br>
	<center id="detailsad">
		<div class="banner" data-id="1074" data-banner-id="1054" data-zone-id="1,2,3,4"><iframe
				src="https://unibet-unibet.bannerflow.com/bf-banners/5a0da8144b0d911cec5dc889.html?cb=636981730599424973&amp;clickpixel=%2F%2F55dacb16e347271ec0d5101b.tracker.bannerflow.com%2Fapi%2Ftr%2Fclick%3Fdata%3D%257B%2522account%2522%253A%2522unibet%2522%252C%2522brand%2522%253A%252255dacb16e347271ec0d5101b%2522%252C%2522placement%2522%253A%25225ade1489890ac04cccdb0e13%2522%252C%2522ad%2522%253A%25225a0da8144b0d911cec5dc88a%2522%252C%2522bannerset%2522%253A%252258a2e4169c0bdc01708d427a%2522%252C%2522banner%2522%253A%25225a0da8144b0d911cec5dc889%2522%252C%2522spotIndexes%2522%253A0%252C%2522bannerIds%2522%253A%25225a0da8144b0d911cec5dc889%2522%257D&amp;targetwindow=_blank&amp;pid=282205&amp;bid=30118&amp;ref=https%3A%2F%2Fb1.trickyrock.com%2Fad.aspx"
				class="bf_animated"
				style="max-width:728pxnone !important; max-height:90pxnone !important;width:728px;height:90px;  background-color: transparent; border: medium none; opacity: 1;"
				scrolling="no" id="bf_16443271218148587" k386ei0vn="" frameborder="0"></iframe></div>
	</center>
	<!--<script type="text/javascript">
    $(document).ready(function(){
        $('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
    });
</script>-->
	<br>
	<div class="torrent_reszletek">
		<div class="torrent_col1">
			<div class="dt">Típus:</div>
			<div class="dd"><a title="Még több ebből a kategóriából"
					href="torrents.php?csoport_listazas=osszes_film">Film</a> &gt; <a
					title="Még több ebből a kategóriából" href="torrents.php?tipus=xvid">SD/EN</a></div>
			<div class="dt">Feltöltve:</div>
			<div class="dd">2013-09-05 12:36:57</div>
			<div class="dt">Feltöltő:</div>
			<div class="dd">
				<span class="feltolto_szin">Anonymous</span>
			</div>
			<div class="dt">Hozzászólás:</div>
			<div class="dd">0 db</div>
		</div>
		<div class="torrent_col2">
			<div class="dt">Seederek:</div>
			<div class="dd"><a onclick="peers('peers', '1365642', 1);" title="Peerlista"
					href="javascript:void(0);">2</a></div>
			<div class="dt">Leecherek:</div>
			<div class="dd"><a onclick="peers('peers', '1365642', 1);" title="Peerlista"
					href="javascript:void(0);">1</a></div>
			<div class="dt">Letöltve:</div>
			<div class="dd">++</div>
			<div class="dt">Sebesség:</div>
			<div class="dd">0 B/s (becsült)</div>
			<div class="dt">Méret:</div>
			<div class="dd">1.69 GiB (1819020964 bájt)</div>
			<div class="dt">Tech info:</div>
			<div class="dd"><a onclick="getNfo('nfocontent', '1365642', 1);" title="Mutat"
					href="javascript:void(0);">Mutat</a></div>
			<div class="dt">Fájlok:</div>
			<div class="dd"><a onclick="files('filelist', '1365642', 1);" title="Fájlok"
					href="javascript:void(0);">1</a></div>
			<!--<div class="dt">Mafab.hu</div>
		<div class="dd"><a onclick="getMafab('mafab', 'tt0120844', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
			<div class="dt">Más verziók:</div>
			<div class="dd"><a onclick="other_versions('otherlist', '0120844', '1365642', 1);" title="Más verziók"
					href="javascript:void(0);">Mutat</a></div>
		</div>
	</div>
	<div style="clear:both;"></div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1365642&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
	<div id="filelist"></div>
	<!-- <div id="mafab"></div>-->
	<div id="otherlist"></div>
	<div id="nfocontent"></div>
	<a name="peers"></a>
	<div id="peers"></div>
	<div id="subtitles"></div>
	<div class="torrent_leiras">
		<table width="100%">
			<tbody>
				<tr>
					<td class="inforbar_img" align="center"><img
							src="https://nc-img.burst.sh/covers/RPev3b57i0mfKBgw" alt="Borító"></td>
					<td class="inforbar_txt">
						<div class="infobar_title"> Star Trek 9. - Űrlázadás<br> (Star Trek: Insurrection)
						</div>
						<table>
							<tbody>
								<tr>
									<td style="vertical-align:top;width:110px;">Megjelenés éve:</td>
									<td>1998</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">DVD Megjelenés:</td>
									<td>2000</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Rendező:</td>
									<td>Jonathan Frakes</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Szereplők:</td>
									<td>Patrick Stewart, Jonathan Frakes, Brent Spiner, LeVar Burton, Michael Dorn</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Ország:</td>
									<td>USA</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Hossz:</td>
									<td>103 perc</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Címkék:</td>
									<td><a style="border-bottom:1px dotted white;"
											href="torrents.php?tags=akci%C3%B3">akció</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=kaland">kaland</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=sci-fi">sci-fi</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=thriller">thriller</a></td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">IMDb értékelés:</td>
									<td>6.3</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">IMDb link:</td>
									<td><a href="https://none.st/?q=https://imdb.com/title/tt0120844"
											target="_blank">https://imdb.com/title/tt0120844</a></td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Egyéb link:</td>
									<td><a href="https://none.st/?q=http://www.port.hu/star_trek_9._-_urlazadas_star_trek_-_insurrection/pls/fi/films.film_page?i_film_id=3626"
											target="_blank">http://www.port.hu/star_trek_9...</a></td>
								</tr>
							</tbody>
						</table>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="torrent_leiras proba42">
		A Szövetség kezdetétől fogva világos volt a Legfőbb Elv: a Csillagflotta egy expedíciója sem avatkozhat be más
		civilizációk természetes fejlődésébe. De Picard most olyan parancsokkal szembesül, melyek ellentmondanak ennek a
		rendeletnek. Ha engedelmeskedik, Ba'ku 600 békés lakosát fogják erőszakkal eltávolítani nagyszerű világukból,
		állítólag milliók érdekében, akik élvezhetik Ba'ku Élet Vize-szerű hatását. Ha Picard nem engedelmeskedik,
		kockáztatja az űrhajóját, a karrierjét, az életét. De Picard számára csak egy választás létezik. Fel kell
		lázadnia a Csillagflotta ellen, és az űrlázadás élére kell állnia, hogy megmentse az Édent.
		<br>
		<br>Hang: angol 5.1
	</div>
	<div class="torrent_leiras">
		Akik eddig megköszönték:<br><br>
		<div id="ncoreKoszonetAjax">
			<a href="profile.php?id=348420">robexx89</a>, <a href="profile.php?id=412001">mdub81</a>, <a
				href="profile.php?id=521417">rapthorock</a>, <a href="profile.php?id=270034">remeset</a>, <a
				href="profile.php?id=747831">deck16</a>, <a href="profile.php?id=711677">Templier</a>, <a
				href="profile.php?id=36003">Apuci7</a>, <a href="profile.php?id=131811">MesM</a>, <a
				href="profile.php?id=655135">szutee</a>, <a href="profile.php?id=790451">Kharma</a>, <a
				href="profile.php?id=289516">minisittes</a>, <a href="profile.php?id=1400527">Neferkat</a>, <a
				href="profile.php?id=1368610">Div6103</a>, <a href="profile.php?id=1332925">Indigo89</a>, <a
				href="profile.php?id=1365098">The1oinogr</a>
			<p align="center">
				<a href="javascript:void(0);" onclick="thanks(1365642);"><img class="g_koszi"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						alt="Megköszönheted a feltöltést"></a>
			</p>
		</div>
	</div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1365642&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
</div>
`

const seriesHtml = `
<div class="fobox_tartalom">
	<div class="torrent_reszletek_cim">Star.Trek.Voyager.S01.iNTERNAL.EXTRAS.DVDRip.x264-MARS</div>
	<div class="torrent_reszletek_konyvjelzo">
		<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);" onclick="konyvjelzo('1972182');"
			title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
		<a style="font-weight:normal;"
			href="torrents.php?action=addnews&amp;id=1972182&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
			title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
		<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=1972182">[bejelentés]</a>
	</div>
	<center>
		<table cellspacing="5" border="0">
			<tbody>
				<tr>
					<td> <a class="fancy_groups" rel="g1972182"
							href="https://nc-img.burst.sh/QoqB738Y0UzCd4kx?24340072"><img
								src="https://nc-img.burst.sh/mDQ42-vPqUVCM4yW?24340072" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g1972182"
							href="https://nc-img.burst.sh/mDQ42-vPqU6fo4yW?24340072"><img
								src="https://nc-img.burst.sh/2lwBDJ70pc-fJ4N8?24340072" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g1972182"
							href="https://nc-img.burst.sh/2lwBDJ70pc7hr4N8?24340072"><img
								src="https://nc-img.burst.sh/Evj4NOg3GI9h0eNx?24340072" border="0"></a></td>
				</tr>
				<tr>
					<td align="center"><em>(640x480)</em></td>
					<td align="center"><em>(640x480)</em></td>
					<td align="center"><em>(640x480)</em></td>
				</tr>
			</tbody>
		</table>
	</center><br>
	<center id="detailsad">
		<div class="banner" data-id="990" data-banner-id="969" data-zone-id="1,2,3,4"><iframe scrolling="no"
				marginwidth="0" noresize="" marginheight="0" vspace="0" hspace="0"
				src="https://rklm1.parom.hu/dyna13/index.php?ref=ncc" width="800" height="100" frameborder="0"></iframe>
		</div>
	</center>
	<!--<script type="text/javascript">
    $(document).ready(function(){
        $('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
    });
</script>-->
	<br>
	<div class="torrent_reszletek">
		<div class="torrent_col1">
			<div class="dt">Típus:</div>
			<div class="dd"><a title="Még több ebből a kategóriából"
					href="torrents.php?csoport_listazas=osszes_sorozat">Sorozat</a> &gt; <a
					title="Még több ebből a kategóriából" href="torrents.php?tipus=xvidser">SD/EN</a></div>
			<div class="dt">Feltöltve:</div>
			<div class="dd">2016-04-11 21:52:35</div>
			<div class="dt">Feltöltő:</div>
			<div class="dd">
				<span class="feltolto_szin">Anonymous</span>
			</div>
			<div class="dt">Hozzászólás:</div>
			<div class="dd">2 db</div>
			<div class="dt">Felirat:</div>
			<div class="dd"><a onclick="getSubs('subtitles', '1972182', 1, 'tt0112178');" title="Feliratok"
					href="javascript:void(0);">Mutat</a></div>
		</div>
		<div class="torrent_col2">
			<div class="dt">Seederek:</div>
			<div class="dd"><a onclick="peers('peers', '1972182', 1);" title="Peerlista"
					href="javascript:void(0);">3</a></div>
			<div class="dt">Leecherek:</div>
			<div class="dd"><a onclick="peers('peers', '1972182', 1);" title="Peerlista"
					href="javascript:void(0);">0</a></div>
			<div class="dt">Letöltve:</div>
			<div class="dd">++</div>
			<div class="dt">Sebesség:</div>
			<div class="dd">0 B/s (becsült)</div>
			<div class="dt">Méret:</div>
			<div class="dd">434.74 MiB (455858033 bájt)</div>
			<div class="dt">NFO:</div>
			<div class="dd"><a onclick="getNfo('nfocontent', '1972182', 1);" title="Mutat"
					href="javascript:void(0);">Mutat</a></div>
			<div class="dt">Fájlok:</div>
			<div class="dd"><a onclick="files('filelist', '1972182', 1);" title="Fájlok"
					href="javascript:void(0);">13</a></div>
			<!--<div class="dt">Mafab.hu</div>
		<div class="dd"><a onclick="getMafab('mafab', 'tt0112178', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
			<div class="dt">Más verziók:</div>
			<div class="dd"><a onclick="other_versions('otherlist', '0112178', '1972182', 1);" title="Más verziók"
					href="javascript:void(0);">Mutat</a></div>
		</div>
	</div>
	<div style="clear:both;"></div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1972182&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
	<div id="filelist"></div>
	<!-- <div id="mafab"></div>-->
	<div id="otherlist"></div>
	<div id="nfocontent"></div>
	<a name="peers"></a>
	<div id="peers"></div>
	<div id="subtitles"></div>
	<div class="torrent_leiras">
		<table width="100%">
			<tbody>
				<tr>
					<td class="inforbar_img" align="center"><img
							src="https://nc-img.burst.sh/covers/lLBo6Kb7hPpSwXvW?24340072" alt="Borító"></td>
					<td class="inforbar_txt">
						<div class="infobar_title"> Star Trek: Voyager </div>
						<table>
							<tbody>
								<tr>
									<td style="vertical-align:top;width:110px;">Megjelenés éve:</td>
									<td>1995</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Rendező:</td>
									<td>David Livingston, Winrich Kolbe</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Szereplők:</td>
									<td>Kate Mulgrew, Robert Beltran, Roxann Dawson, Robert Duncan McNeill, Ethan
										Phillips</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Ország:</td>
									<td>USA</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Hossz:</td>
									<td>45 perc</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Címkék:</td>
									<td><a style="border-bottom:1px dotted white;"
											href="torrents.php?tags=akci%C3%B3">akció</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=kaland">kaland</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=sci-fi">sci-fi</a></td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">IMDb értékelés:</td>
									<td>7.7</td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">IMDb link:</td>
									<td><a href="https://none.st/?q=https://imdb.com/title/tt0112178"
											target="_blank">https://imdb.com/title/tt0112178</a></td>
								</tr>
								<tr>
									<td style="vertical-align:top;width:110px;">Egyéb link:</td>
									<td><a href="https://none.st/?q=http://www.port.hu/star_trek:_voyager/pls/w/films.film_page?i_film_id=26841"
											target="_blank">http://www.port.hu/star_trek:_...</a></td>
								</tr>
							</tbody>
						</table>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="torrent_leiras proba42">
		Eredeti release!
		<br>
		<br>
		<br>
		<br>
	</div>
	<div class="torrent_leiras">
		Akik eddig megköszönték:<br><br>
		<div id="ncoreKoszonetAjax">
			<p align="center">
				<a href="javascript:void(0);" onclick="thanks(1972182);"><img class="g_koszi"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						alt="Megköszönheted a feltöltést"></a>
			</p>
		</div>
	</div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1972182&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
</div>
`

const gameHtml = `
<div id="main_tartalom">
	<script type="text/javascript">
		function fads() {
			return '<iframe src="http://www.google.hu/intl/hu/privacy/" style="width:650px;height:70px;overflow:hidden;border:none;"></iframe>';
		}
		$(document).ready(function () {
			$('a.fancy_groups').fancybox({
				'type': 'image'
			});
			$('a.fancy_trailer').fancybox({
				'autoScale': false,
				'width': 655,
				'height': 500,
				'type': 'iframe'
			});
		});
		function toggle_modpanel(obj) {
			if (document.getElementById('modpanel_div').style.display == '') {
				document.getElementById('modpanel_div').style.display = 'none';
				obj.innerHTML = "Moderátor panel mutatása";
			} else {
				document.getElementById('modpanel_div').style.display = '';
				obj.innerHTML = "Moderátor panel elrejtése";
			}
		}
		function komment_torol(id, id2, uni) {
			if (confirm('Biztosan törlöd ezt a kommentet?')) {
				window.location.href = 'torrents.php?action=comments&a=torol&id=' + id + '&next_comment_id=' + id2 + '&getunique=' + uni;
			} else {
				return false;
			}
		}
		function komment_all_torol(id, uni) {
			if (confirm('Biztosan törlöd az ÖSSZES kommentet?')) {
				window.location.href = 'torrents.php?action=comments&a=torol_all&id=' + id + '&getunique=' + uni;
			} else {
				return false;
			}
		}
		function nCoreKep(cim, meretek) {
			var url = cim;
			var meret = meretek.split("x", 2);
			var width = ((parseInt(meret[0]) + 25) > 455 ? (parseInt(meret[0]) + 25) : 455);
			var height = parseInt(meret[1]) + 170;
			window.open('/showimage.php?link=' + url, '_blank', 'toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=' + width + ', height=' + height);
		}
		function nCoreTrailer(cim) {
			var url = cim;
			window.open('/showtrailer.php?link=' + url, '_blank', 'toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=675, height=520');
		}
	</script>
	<div class="fobox_all" id="details1">
		<div class="fobox_fej_t"></div>
		<div class="fobox_fej">Torrent adatai</div>
		<div class="fobox_fej_b"></div>
		<div class="fobox_tartalom_t"></div>
		<div class="fobox_tartalom">
			<div class="torrent_reszletek_cim">Baldur's Gate - Enhanced Edition v2.6.5.0.45309 GOG</div>
			<div class="torrent_reszletek_konyvjelzo">
				<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);"
					onclick="konyvjelzo('3176798');" title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
				<a style="font-weight:normal;"
					href="torrents.php?action=addnews&amp;id=3176798&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
					title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
				<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=3176798">[bejelentés]</a>
			</div>
			<center>
				<table cellspacing="5" border="0">
					<tbody>
						<tr>
							<td> <a class="fancy_groups" rel="g3176798"
									href="https://nc-img.burst.sh/Jl3X9v9W6fyCD_7R?26990706"><img
										src="https://nc-img.burst.sh/y7D4rJ-oPT5C84vV?26990706" border="0"></a>
							</td>
							<td> <a class="fancy_groups" rel="g3176798"
									href="https://nc-img.burst.sh/y7D4rJ-oPTbf74vV?26990706"><img
										src="https://nc-img.burst.sh/6ZV_31yAqiGfDepM?26990706" border="0"></a>
							</td>
							<td> <a class="fancy_groups" rel="g3176798"
									href="https://nc-img.burst.sh/6ZV_31yAqiJhwepM?26990706"><img
										src="https://nc-img.burst.sh/J7YBQElrjcGhqXdn?26990706" border="0"></a>
							</td>
						</tr>
						<tr>
							<td align="center"><em>(1980x1215)</em></td>
							<td align="center"><em>(1920x1080)</em></td>
							<td align="center"><em>(1920x1080)</em></td>
						</tr>
					</tbody>
				</table>
			</center><br>
			<center id="detailsad">
				<div class="banner" data-id="1074" data-banner-id="1054" data-zone-id="1,2,3,4"><iframe
						src="https://unibet-unibet.bannerflow.com/bf-banners/5a0da8144b0d911cec5dc889.html?cb=636981730599424973&amp;clickpixel=%2F%2F55dacb16e347271ec0d5101b.tracker.bannerflow.com%2Fapi%2Ftr%2Fclick%3Fdata%3D%257B%2522account%2522%253A%2522unibet%2522%252C%2522brand%2522%253A%252255dacb16e347271ec0d5101b%2522%252C%2522placement%2522%253A%25225ade1489890ac04cccdb0e13%2522%252C%2522ad%2522%253A%25225a0da8144b0d911cec5dc88a%2522%252C%2522bannerset%2522%253A%252258a2e4169c0bdc01708d427a%2522%252C%2522banner%2522%253A%25225a0da8144b0d911cec5dc889%2522%252C%2522spotIndexes%2522%253A0%252C%2522bannerIds%2522%253A%25225a0da8144b0d911cec5dc889%2522%257D&amp;targetwindow=_blank&amp;pid=282205&amp;bid=30118&amp;ref=https%3A%2F%2Fb1.trickyrock.com%2Fad.aspx"
						class="bf_animated"
						style="max-width:728pxnone !important; max-height:90pxnone !important;width:728px;height:90px;  background-color: transparent; border: medium none; opacity: 1;"
						scrolling="no" id="bf_16443271218148587" srirgynu2="" frameborder="0"></iframe></div>
			</center>
			<!--<script type="text/javascript">
		$(document).ready(function(){
			$('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
		});
	</script>-->
			<br>
			<div class="torrent_reszletek">
				<div class="torrent_col1">
					<div class="dt">Típus:</div>
					<div class="dd"><a title="Még több ebből a kategóriából"
							href="torrents.php?csoport_listazas=osszes_jatek">Játék</a> &gt; <a
							title="Még több ebből a kategóriából" href="torrents.php?tipus=game_rip">PC/RIP</a></div>
					<div class="dt">Feltöltve:</div>
					<div class="dd">2021-04-26 15:06:45</div>
					<div class="dt">Feltöltő:</div>
					<div class="dd">
						<span class="feltolto_szin">Anonymous</span>
					</div>
					<div class="dt">Hozzászólás:</div>
					<div class="dd">0 db</div>
				</div>
				<div class="torrent_col2">
					<div class="dt">Seederek:</div>
					<div class="dd"><a onclick="peers('peers', '3176798', 1);" title="Peerlista"
							href="javascript:void(0);">15</a></div>
					<div class="dt">Leecherek:</div>
					<div class="dd"><a onclick="peers('peers', '3176798', 1);" title="Peerlista"
							href="javascript:void(0);">1</a></div>
					<div class="dt">Letöltve:</div>
					<div class="dd">++</div>
					<div class="dt">Sebesség:</div>
					<div class="dd">0 B/s (becsült)</div>
					<div class="dt">Méret:</div>
					<div class="dd">5.15 GiB (5531492728 bájt)</div>
					<div class="dt">Tech info:</div>
					<div class="dd"><a onclick="getNfo('nfocontent', '3176798', 1);" title="Mutat"
							href="javascript:void(0);">Mutat</a></div>
					<div class="dt">Fájlok:</div>
					<div class="dd"><a onclick="files('filelist', '3176798', 1);" title="Fájlok"
							href="javascript:void(0);">152</a></div>
					<!--<div class="dt">Mafab.hu</div>
			<div class="dd"><a onclick="getMafab('mafab', 'tt', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
				</div>
			</div>
			<div style="clear:both;"></div>
			<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
					href="torrents.php?action=download&amp;id=3176798&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
					letöltése</a></div>
			<div id="filelist"></div>
			<!-- <div id="mafab"></div>-->
			<div id="otherlist"></div>
			<div id="nfocontent"></div>
			<a name="peers"></a>
			<div id="peers"></div>
			<div id="subtitles"></div>
			<div class="torrent_leiras">
				<table width="100%">
					<tbody>
						<tr>
							<td class="inforbar_img" align="center"><img
									src="https://nc-img.burst.sh/covers/jg_mpnRviRwfle8z?26990706" alt="Borító">
							</td>
						</tr>
					</tbody>
				</table>
			</div>
			<div class="torrent_leiras proba42">
				<span style="font-style:italic">CÍM: Baldur's Gate: Enhanced Edition
					<br>MŰFAJ: Kaland, Szerepjáték
					<br>FEJLESZTŐ: Beamdog
					<br>KIADÓ: Beamdog
					<br>MEGJELENÉS DÁTUMA: 2013. jan. 16.
					<br>
					<br><span style="text-decoration:underline">Description:</span>
					<br>
					<br>Baldur's Gate: Enhanced Edition includes the classic Baldur's Gate: The Original Saga.
					<br>
					<br>Gather Your Party
					<br>Baldur's Gate: Enhanced Edition is a story-driven 90s RPG set in the world of Dungeons &amp;
					Dragons.
					<br>
					<br>Customize your hero, recruit a party of brave allies, and explore the Sword Coast in your search
					for adventure, profit… and the truth.
					<br>75+ Hours of Adventure
					<br>The Enhanced Edition contains over 75 hours of gameplay, including the original campaign, the
					classic Sword Coast expansion, plus brand new challenges in the Black Pits arena!
					<br>
					<br>Classic Campaign: The Original Baldur’s Gate Adventure
					<br>Expansion: Tales of the Sword Coast expansion
					<br>New Challenges: The Black Pits, arena style battles
					<br>New Difficulty Setting: Story Mode allows players to focus on story and exploration, rather than
					combat and survival
					<br>Paid DLC Expansion Available: Siege of Dragonspear is a brand new chapter in the Baldur's Gate
					saga!
					<br>Epic Characters
					<br>
					<br>11 Playable Classes plus dozens of subclasses
					<br>Recruit Classic Characters like Minsc and his brave hamster, Boo!
					<br>3 New Recruitable Heroes: Neera the Wild Mage, Dorn Il-Khan the Blackguard, and Rasaad yn Bashir
					the Monk
					<br>New player voice sets to customize your hero
					<br>Story-driven gameplay means character choices matter
					<br>
					<br>Classic Gameplay
					<br>2-D isometric graphics
					<br>Real-time-with-pause combat
					<br>Adapts 2nd Edition Dungeons &amp; Dragons Rules
					<br>
					<br>Enhanced for Modern Platforms
					<br>Over 400 improvements to the original game
					<br>Native support for high-resolution widescreen displays
					<br>The 1998 Classic, enhanced for modern Windows, macOS and Linux players!
					<br>Story-Rich Gaming Experience
					<br>Forced to leave your home under mysterious circumstances, you find yourself drawn into a
					conflict that has the Sword Coast on the brink of war.
					<br>
					<br>Your view of the world has been limited to the heavily fortified walls of Candlekeep. Your
					foster father, Gorion, has done everything in his power to protect you, and keep you out of harm’s
					way. All that is about to change...
					<br>
					<br><span style="text-decoration:underline">RENDSZERKÖVETELMÉNYEK:</span>
					<br>
					<br><span style="text-decoration:underline">MINIMUM:</span>
					<br>OS:Windows 7, 8.1, 10, 11 64 bit
					<br>Processor: Dual Core Processor
					<br>Memory: 1 GB RAM
					<br>Graphics: OpenGL 2.0 compatible
					<br>Hard Drive: 5 GB available space
					<br>
					<br><span style="text-decoration:underline">Goodies Included:</span>
					<br>
					<br>baldurs_gate_1_ee_mastering_melee_and_magic
					<br>baldurs_gate_1_ee_sword_coast_survival_guide
					<br>baldurs_gate_avatars
					<br>baldurs_gate_series_wallpapers
					<br>bg_ee_soundtrack_flac
					<br>bg_ee_soundtrack_mp3
					<br>bg_sod_adventurers_guide
					<br>bg_sod_field_guide
					<br>bg_sod_soundtrack_flac
					<br>bg_sod_soundtrack_mp3
					<br>bg_sod_survival_guide
					<br>siege_of_dragonspear_artworks
					<br>
					<br><span style="text-decoration:underline">DLCs:</span>
					<br>
					<br>Baldur's Gate: Siege of Dragonspear
					<br>Baldur's Gate: Faces of Good and Evil
					<br>
					<br><a href="https://none.st/?q=https://www.gog.com/game/baldurs_gate_enhanced_edition"
						target="_blank" class="bb-url">https://www.gog.com/game/baldurs_gate_enhanced_edition</a>
					<br>
					<br><span style="text-decoration:underline">Telepítés:</span>
					<br>
					<br>- A setup_baldurs_gate_enhanced_edition_2.6.5.0_(64bit)_(45309).exe nevű fájllal elindítod a
					telepítőt.
					<br>- Kiválasztod a kívánt komponenseket (nyelv stb.)
					<br>- DLCk telepítése
					<br>- A telepítés végeztével az asztalról elindítod a játékot
					<br>
					<br><span style="text-decoration:underline">Gameplay:</span>
					<br>
					<div style="margin-top:5px;text-align:left;"><a href="javascript:void(0);"
							onclick="if (this.innerHTML=='Spoiler megjelenítése') { this.innerHTML='Spoiler elrejtése';this.parentNode.nextSibling.childNodes[0].style.display='none';this.parentNode.nextSibling.childNodes[1].style.display='block';} else {this.innerHTML='Spoiler megjelenítése';this.parentNode.nextSibling.childNodes[0].style.display='block';this.parentNode.nextSibling.childNodes[1].style.display='none';}">Spoiler
							megjelenítése</a></div>
					<div class="bb-quote">
						<div style="display: block;">... spoiler ...</div>
						<div style="display: none;">
							<div class="js-video widescreen" align="center"><iframe
									src="https://www.youtube.com/embed/S2RheHL7YOI" allowfullscreen="" width="610"
									height="502" frameborder="0"></iframe></div>
						</div>
					</div>
				</span>
			</div>
			<div class="torrent_leiras">
				Akik eddig megköszönték:<br><br>
				<div id="ncoreKoszonetAjax">
					<p align="center">
						<a href="javascript:void(0);" onclick="thanks(3176798);"><img class="g_koszi"
								src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
								alt="Megköszönheted a feltöltést"></a>
					</p>
				</div>
			</div>
			<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
					href="torrents.php?action=download&amp;id=3176798&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
					letöltése</a></div>
		</div>
		<div class="fobox_lab"></div>
	</div>
	<div align="center">
		<div style="width:917px;height:13px;"></div>
	</div>
	<a name="comments"></a>
	<div class="fobox_all" id="details2">
		<div class="fobox_fej_t"></div>
		<div class="fobox_fej">Hozzászólások</div>
		<div class="fobox_fej_b"></div>
		<div class="fobox_tartalom_t"></div>
		<div class="fobox_tartalom">
			<div id="hozzaszolasok">
				<div class="teljes_oldalas">Ehhez a torrenthez még senki nem írt hozzászólást!</div>
			</div>
		</div>
		<div class="fobox_lab"></div>
	</div>
	<div align="center">
		<div style="width:917px;height:13px;"></div>
	</div>
	<div class="commentbox_all">
		<div class="commentbox_fej">Szólj hozzá te is!</div>
		<div class="commentbox_tartalom">
			<a name="reply"></a>
			<form action="torrents.php?action=details&amp;id=3176798#comment_pre" method="post" name="hozzaszol"
				target="_self">
				<input type="hidden" name="getUnique" value="8ffe570cf24f14b749ee44e32587c118">

				<div class="hsz_iconok">
					<table class="hsz_iconok_table">
						<tbody>
							<tr>
								<td class="hsz_width1">
									<div class="hsz_bb_vastag"><a href="javascript: AddBB('szoveg', 'bold')"
											title="Vastag betű"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_dolt"><a href="javascript: AddBB('szoveg', 'italic')"
											title="Dőlt betű"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_alahuzott"><a href="javascript: AddBB('szoveg', 'underline')"
											title="Aláhúzott betű"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
								</td>
								<td class="hsz_width2">
									<select id="sizes" name="size" size="1">
										<option value="8pt">8</option>
										<option value="9pt">9</option>
										<option selected="selected" value="10pt">10</option>
										<option value="12pt">12</option>
										<option value="14pt">14</option>
										<option value="16pt">16</option>
										<option value="18pt">18</option>
										<option value="20pt">20</option>
										<option value="22pt">22</option>
										<option value="24pt">24</option>
										<option value="30pt">30</option>
										<option value="36pt">36</option>
										<option value="48pt">48</option>
										<option value="72pt">72</option>
									</select>
								</td>
								<td class="hsz_width3">
									<div class="hsz_bb_meret"><a href="javascript: AddBB('szoveg', 'size')"
											title="Betűméret"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
								</td>
								<td class="hsz_width4">
									<select id="colors" name="color" size="1">
										<option selected="selected" value="black">Fekete</option>
										<option value="white">Fehér</option>
										<option value="green">Zöld</option>
										<option value="maroon">Gesztenyebarna</option>
										<option value="olive">Olivazöld</option>
										<option value="navy">Mélykék</option>
										<option value="purple">Lila</option>
										<option value="gray">Szürke</option>
										<option value="yellow">Sárga</option>
										<option value="lime">Lime</option>
										<option value="aqua">Cián</option>
										<option value="fuchsia">Ciklámen</option>
										<option value="silver">Ezüst</option>
										<option value="red">Piros</option>
										<option value="blue">Kék</option>
										<option value="teal">Pávakék</option>
									</select>
								</td>
								<td class="hsz_width5">
									<div class="hsz_bb_szin"><a href="javascript: AddBB('szoveg', 'color')"
											title="Betűszín"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_smilie"><a target="_blank" href="wiki.php?action=read&amp;id=403"
											title="Smile beszúrása"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_kep"><a href="javascript: AddBB('szoveg', 'picture')"
											title="Kép beszúrása"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_imgw"><a href="javascript: AddBB('szoveg', 'imgw')"
											title="Nagyméretű kép beszúrása"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_link"><a href="javascript: AddBB('szoveg', 'url')"
											title="Link beszúrása"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_spoiler"><a href="javascript: AddBB('szoveg', 'spoiler')"
											title="Spoiler, rejtett szöveg beszúrása"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_felsorolas"><a href="javascript: AddBB('szoveg', 'list')"
											title="Felsorolás"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_idezet"><a href="javascript: AddBB('szoveg', 'quote')"
											title="Idézet"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
									<div class="hsz_bb_kerdes"><a target="_blank" href="wiki.php?action=read&amp;id=393"
											title="Kérdés"><img class="bb_link"
												src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"></a>
									</div>
								</td>
							</tr>
						</tbody>
					</table>
				</div>
				<div class="hsz_textarea">
					<textarea class="textarea_fancy" id="szoveg" name="szoveg"
						onbeforedeactivate="window.txtRange = document.selection.createRange();"
						tabindex="1"></textarea>
				</div>
				<br><br>
				<center>
					<input name="g_mehet" type="image" value="submit"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						class="g_mehet">
					&nbsp;&nbsp;&nbsp;
					<input name="g_elonezet" type="image" value="submit"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						class="g_elonezet">
				</center>
			</form>
		</div>
		<div class="commentbox_lab"></div>
	</div>
</div>
`

const musicHtml = `
<div class="fobox_tartalom">
	<div class="torrent_reszletek_cim">Linkin Park - Greatest Hits</div>
	<div class="torrent_reszletek_konyvjelzo">
		<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);" onclick="konyvjelzo('1475272');"
			title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
		<a style="font-weight:normal;"
			href="torrents.php?action=addnews&amp;id=1475272&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
			title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
		<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=1475272">[bejelentés]</a>
	</div>
	<br>
	<center id="detailsad">
		<div class="banner" data-id="1070" data-banner-id="1050" data-zone-id="1,2,3,4"><iframe
				src="https://unibet-unibet.bannerflow.com/bf-banners/611a249006739a7cdcd7b670.html?cb=637726710070815792&amp;clickpixel=%2F%2F55dacb16e347271ec0d5101b.tracker.bannerflow.com%2Fapi%2Ftr%2Fclick%3Fdata%3D%257B%2522account%2522%253A%2522unibet%2522%252C%2522brand%2522%253A%252255dacb16e347271ec0d5101b%2522%252C%2522placement%2522%253A%25225d1cc5779aa5c23584d10826%2522%252C%2522ad%2522%253A%25225d1cc5359aa5c23584d1080e%2522%252C%2522bannerset%2522%253A%2522610c058b747833d603aedefa%2522%252C%2522banner%2522%253A%2522611a249006739a7cdcd7b670%2522%252C%2522spotIndexes%2522%253A0%252C%2522bannerIds%2522%253A%2522611a249006739a7cdcd7b670%2522%257D&amp;targetwindow=_blank&amp;pid=282205&amp;bid=34111&amp;ref=https%3A%2F%2Fb1.trickyrock.com%2Fad.aspx"
				class="bf_animated"
				style="max-width:728pxnone !important; max-height:90pxnone !important;width:728px;height:90px;  background-color: transparent; border: medium none; opacity: 1;"
				scrolling="no" id="bf_16443269015075793" hmx9wf4m3="" frameborder="0"></iframe></div>
	</center>
	<!--<script type="text/javascript">
$(document).ready(function(){
	$('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
});
</script>-->
	<br>
	<div class="torrent_reszletek">
		<div class="torrent_col1">
			<div class="dt">Típus:</div>
			<div class="dd"><a title="Még több ebből a kategóriából"
					href="torrents.php?csoport_listazas=osszes_zene">Zene</a> &gt; <a
					title="Még több ebből a kategóriából" href="torrents.php?tipus=mp3">MP3/EN</a></div>
			<div class="dt">Feltöltve:</div>
			<div class="dd">2014-01-25 18:41:23</div>
			<div class="dt">Feltöltő:</div>
			<div class="dd">
				<span class="feltolto_szin">Anonymous</span>
			</div>
			<div class="dt">Hozzászólás:</div>
			<div class="dd">1 db</div>
		</div>
		<div class="torrent_col2">
			<div class="dt">Seederek:</div>
			<div class="dd"><a onclick="peers('peers', '1475272', 1);" title="Peerlista"
					href="javascript:void(0);">230</a></div>
			<div class="dt">Leecherek:</div>
			<div class="dd"><a onclick="peers('peers', '1475272', 1);" title="Peerlista"
					href="javascript:void(0);">2</a></div>
			<div class="dt">Letöltve:</div>
			<div class="dd">+++++</div>
			<div class="dt">Sebesség:</div>
			<div class="dd">99.78 KiB/s (becsült)</div>
			<div class="dt">Méret:</div>
			<div class="dd">366.12 MiB (383904903 bájt)</div>
			<div class="dt">Tech info:</div>
			<div class="dd"><a onclick="getNfo('nfocontent', '1475272', 1);" title="Mutat"
					href="javascript:void(0);">Mutat</a></div>
			<div class="dt">Fájlok:</div>
			<div class="dd"><a onclick="files('filelist', '1475272', 1);" title="Fájlok"
					href="javascript:void(0);">46</a></div>
			<!--<div class="dt">Mafab.hu</div>
	<div class="dd"><a onclick="getMafab('mafab', 'tt', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
		</div>
	</div>
	<div style="clear:both;"></div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1475272&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
	<div id="filelist"></div>
	<!-- <div id="mafab"></div>-->
	<div id="otherlist"></div>
	<div id="nfocontent"></div>
	<a name="peers"></a>
	<div id="peers"></div>
	<div id="subtitles"></div>
	<div class="torrent_leiras">
		<table width="100%">
			<tbody>
				<tr>
					<td class="inforbar_img" align="center"><img
							src="https://nc-img.burst.sh/covers/DQ_2ZqxWT69io_yW" alt="Borító"></td>
					<td class="inforbar_txt">
						<table>
							<tbody>
								<tr>
									<td style="vertical-align:top;width:110px;">Címkék:</td>
									<td><a style="border-bottom:1px dotted white;"
											href="torrents.php?tags=alternative">alternative</a>, <a
											style="border-bottom:1px dotted white;"
											href="torrents.php?tags=metal">metal</a></td>
								</tr>
							</tbody>
						</table>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="torrent_leiras proba42">
		<span style="text-decoration:underline">Előadó</span>: Linkin Park
		<br><span style="text-decoration:underline">Album címe</span>: Greatest Hits
		<br><span style="text-decoration:underline">Megjelenés éve</span>: 2014
		<br><a
			href="https://none.st/?q=http://www.ultramusicstore.com/linkin-park-greatest-hits-2cd-set-digipack-p-4221.html"
			target="_blank"
			class="bb-url">http://www.ultramusicstore.com/linkin-park-greatest-hits-2cd-set-digipack-p-4221.html</a>
		<br>
		<br>Tracklist:
		<br>
		<br>CD-1
		<br>01. Burn It Down
		<br>02. New Divide
		<br>03. What I've Done
		<br>04. The Catalyst
		<br>05. Lost In The Echo
		<br>06. Bleed It Out
		<br>07. Numb
		<br>08. One Step Closer
		<br>09. Leave Out All The Rest
		<br>10. Somewhere I Belong
		<br>11. Given Up
		<br>12. Points Of Authority
		<br>13. No More Sorrow
		<br>14. Breaking The Habit
		<br>15. Runaway
		<br>16. Lying From You
		<br>17. In The End
		<br>18. Shadow Of The Day
		<br>19. A Place For My Head
		<br>20. Faint
		<br>21. Papercut
		<br>22. Crawling
		<br>23. From The Inside
		<br>
		<br>CD-2
		<br>01. In My Remains
		<br>02. Waiting For The End
		<br>03. Lies Greed Misery
		<br>04. Burning In The Skies
		<br>05. Castle Of Glass
		<br>06. Iridescent
		<br>07. Blackout
		<br>08. Foreword
		<br>09. Don'T Stay
		<br>10. With You
		<br>11. No Roads Left
		<br>12. Easier To Run
		<br>13. Qwerty
		<br>14. Pushing Me Away
		<br>15. And One
		<br>16. Hit The Floor
		<br>17. Valentines Day
		<br>18. By Myself
		<br>19. Figure.09
		<br>20. Forgotten
		<br>21. In Pieces
		<br>22. In Between
		<br>23. My December
	</div>
	<div class="torrent_leiras">
		Akik eddig megköszönték:<br><br>
		<div id="ncoreKoszonetAjax">
			<p align="center">
				<a href="javascript:void(0);" onclick="thanks(1475272);"><img class="g_koszi"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						alt="Megköszönheted a feltöltést"></a>
			</p>
		</div>
	</div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1475272&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
</div>
`

const appHtml = `
<div class="fobox_tartalom">
	<div class="torrent_reszletek_cim">Total Commander v10.0 HUN-ENG x86-x64</div>
	<div class="torrent_reszletek_konyvjelzo">
		<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);" onclick="konyvjelzo('3194806');"
			title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
		<a style="font-weight:normal;"
			href="torrents.php?action=addnews&amp;id=3194806&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
			title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
		<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=3194806">[bejelentés]</a>
	</div>
	<center>
		<table cellspacing="5" border="0">
			<tbody>
				<tr>
					<td> <a class="fancy_groups" rel="g3194806"
							href="https://nc-img.burst.sh/jvw_ZoJZkSrCzBNl?27057497"><img
								src="https://nc-img.burst.sh/N-jB-ONG7C-CZ_0v?27057497" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g3194806"
							href="https://nc-img.burst.sh/N-jB-ONG7CRfz_0v?27057497"><img
								src="https://nc-img.burst.sh/AL9XkMlLKCwf9XZl?27057497" border="0"></a></td>
					<td> <a class="fancy_groups" rel="g3194806"
							href="https://nc-img.burst.sh/AL9XkMlLKCLhlXZl?27057497"><img
								src="https://nc-img.burst.sh/kRPBvElwLTQhY4gw?27057497" border="0"></a></td>
				</tr>
				<tr>
					<td align="center"><em>(328x303)</em></td>
					<td align="center"><em>(1484x810)</em></td>
					<td align="center"><em>(499x445)</em></td>
				</tr>
			</tbody>
		</table>
	</center><br>
	<center id="detailsad">
		<div class="banner" data-id="1074" data-banner-id="1054" data-zone-id="1,2,3,4"><iframe
				src="https://unibet-unibet.bannerflow.com/bf-banners/5a0da8144b0d911cec5dc889.html?cb=636981730599424973&amp;clickpixel=%2F%2F55dacb16e347271ec0d5101b.tracker.bannerflow.com%2Fapi%2Ftr%2Fclick%3Fdata%3D%257B%2522account%2522%253A%2522unibet%2522%252C%2522brand%2522%253A%252255dacb16e347271ec0d5101b%2522%252C%2522placement%2522%253A%25225ade1489890ac04cccdb0e13%2522%252C%2522ad%2522%253A%25225a0da8144b0d911cec5dc88a%2522%252C%2522bannerset%2522%253A%252258a2e4169c0bdc01708d427a%2522%252C%2522banner%2522%253A%25225a0da8144b0d911cec5dc889%2522%252C%2522spotIndexes%2522%253A0%252C%2522bannerIds%2522%253A%25225a0da8144b0d911cec5dc889%2522%257D&amp;targetwindow=_blank&amp;pid=282205&amp;bid=30118&amp;ref=https%3A%2F%2Fb1.trickyrock.com%2Fad.aspx"
				class="bf_animated"
				style="max-width:728pxnone !important; max-height:90pxnone !important;width:728px;height:90px;  background-color: transparent; border: medium none; opacity: 1;"
				scrolling="no" id="bf_16443271218148587" hsqr6po8d="" frameborder="0"></iframe></div>
	</center>
	<!--<script type="text/javascript">
$(document).ready(function(){
	$('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
});
</script>-->
	<br>
	<div class="torrent_reszletek">
		<div class="torrent_col1">
			<div class="dt">Típus:</div>
			<div class="dd"><a title="Még több ebből a kategóriából"
					href="torrents.php?csoport_listazas=osszes_program">Program</a> &gt; <a
					title="Még több ebből a kategóriából" href="torrents.php?tipus=misc">Prog/RIP</a></div>
			<div class="dt">Feltöltve:</div>
			<div class="dd">2021-06-12 00:17:57</div>
			<div class="dt">Feltöltő:</div>
			<div class="dd">
				<span class="feltolto_szin">Anonymous</span>
			</div>
			<div class="dt">Hozzászólás:</div>
			<div class="dd">21 db</div>
		</div>
		<div class="torrent_col2">
			<div class="dt">Seederek:</div>
			<div class="dd"><a onclick="peers('peers', '3194806', 1);" title="Peerlista"
					href="javascript:void(0);">1396</a></div>
			<div class="dt">Leecherek:</div>
			<div class="dd"><a onclick="peers('peers', '3194806', 1);" title="Peerlista"
					href="javascript:void(0);">36</a></div>
			<div class="dt">Letöltve:</div>
			<div class="dd">+++++</div>
			<div class="dt">Sebesség:</div>
			<div class="dd">4.51 KiB/s (becsült)</div>
			<div class="dt">Méret:</div>
			<div class="dd">10.41 MiB (10911488 bájt)</div>
			<div class="dt">Tech info:</div>
			<div class="dd"><a onclick="getNfo('nfocontent', '3194806', 1);" title="Mutat"
					href="javascript:void(0);">Mutat</a></div>
			<div class="dt">Fájlok:</div>
			<div class="dd"><a onclick="files('filelist', '3194806', 1);" title="Fájlok"
					href="javascript:void(0);">3</a></div>
			<!--<div class="dt">Mafab.hu</div>
	<div class="dd"><a onclick="getMafab('mafab', 'tt', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
		</div>
	</div>
	<div style="clear:both;"></div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=3194806&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
	<div id="filelist"></div>
	<!-- <div id="mafab"></div>-->
	<div id="otherlist"></div>
	<div id="nfocontent"></div>
	<a name="peers"></a>
	<div id="peers"></div>
	<div id="subtitles"></div>
	<div class="torrent_leiras">
		<table width="100%">
			<tbody>
				<tr>
					<td class="inforbar_img" align="center"><img
							src="https://nc-img.burst.sh/covers/PxBzREgbSMlT-BbM?27057497" alt="Borító"></td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="torrent_leiras proba42">

		<br><span style="font-size:14pt"><span class="highlight">Total Commander 10.0 végleges verzió kiadásra
				került.</span>
			<br></span>
		<br><span style="font-size:12pt">A legfőbb újdonságok:
			<br></span>
		<br><span class="highlight">Általános:</span>
		<br>- A MacOS-el létrehozott ExFAT flash meghajtókon a könyvtárak mostantól a TC-ben is megjelennek.
		<br>- Az "AppData" rejtett könyvtár a c:&nbsp;&nbsp;Users&nbsp;&nbsp;-ban mindig látszik, akkor is, ha a rejtett
		fájlok le vannak tiltva
		<br>- A regisztrációs kulcsfájl automatikus telepítése dupla kattintással
		<br>- Ha fájlokat bont ki belső vagy plug-in segítségével, átmásolja a "Zone Transfer" adatot az archívumból a
		kicsomagolt fájlokba, akárcsak az Explorer
		<br>- A "Windows Hello" használata a Windows 10 rendszeren a fő jelszó biztonságos tárolásához, amíg a TC le nem
		zárul
		<br>- Ha nincsenek fülek, az egér görgető használható a mappák közötti görgetéshez (mint Pl. a FireFoxban)
		<br>- A miniatűrök betöltése EXIF-adatokból (előnézeti kép) JPG-ből és különféle fényképezőgép RAW
		formátumokból: Canon RAW 1-3, DNG, Fuji, Nikon, Olympus, Panasonic, Pentax, Sony
		<br>- "!" jel, ha az aktuális mappa különbözik a zárolt mappától azokon a füleken, amelyeket a
		könyvtárváltoztatások engedélyezve vannak
		<br>- a parancssori utasítás csak akkor hajtódik végre, ha a fókuszban van
		<br>- Ha ideiglenes fájlok vannak nyitva, bezáráskor mutatja az első 3 ilyen fájl nevét
		<br>- Fájl behúzás másik programból: új dialógus a jobb felülírás, másolás kezelésére
		<br>- Ha a felhasználó nem tudja menteni a beállításokat (például a Konfiguráció menüben vagy a fő
		beállításokban), felajánlja az ini fájlok másolását a felhasználói profilba
		<br>- Megadható különböző ini tárolási hely az egyes bővítményekhez: wincmd.ini [ReplaceIniLocation]
		<br><span class="highlight">Nézőke:</span>
		<br>- Új DirectShow médialejátszó a Nézőkében, felváltva a régit
		<br>- Teljes képernyős lejátszás, dupla kattintással
		<br>- Gombokkal léptetés előre, hátra 10,60 másodperccel
		<br><span class="highlight">Szinkronizáció:
			<br></span>- Mentett reguláris kifejezések és keresések fejlettebb használata
		<br>- Szinkronizáció: Újabb helyi menük a tulajdonságok megjelenítéséhez és szerkesztéséhez
		<br><span class="highlight">Tartalom összehasonlítás:
			<br></span>- Extra lista megjelenítés a kurzor aktuális helyzetének két sorával a könnyebb karakter-karakter
		összehasonlítás érdekében
		<br>- Párbeszédpanel megjelenítése, ha a nagy szöveges puffer beillesztése több mint 2 másodpercet vesz igénybe,
		lehetővé teszi a beillesztés megszakítását
		<br><span class="highlight">FTP:</span>
		<br>- Támogatja az OpenSSL 1.1 vagy újabb verziót.
		<br>
		<br><span class="highlight">Hivatalos oldalak:</span>
		<br><a href="https://none.st/?q=https://www.ghisler.com/" target="_blank"
			class="bb-url">https://www.ghisler.com/</a>
		<br><a href="https://none.st/?q=https://www.totalcommander.hu/" target="_blank"
			class="bb-url">https://www.totalcommander.hu/</a>
		<br>
		<br><span class="highlight">Támogatott operációs rendszerek:</span>
		<br>Windows® 95/98 / ME / NT / 2000 / XP / Vista / 7/8 / 8.1 / 10 és a Windows® 3.1
		<br>
		<br><span class="highlight">Telepítés:</span>
		<br>Kiválasztod a neked megfelelő telepítőt, vagy mindkettőt (ha Windows 64-ed van).
		<br>Frissítheted a meglévőt is.
		<br>A "Crack" mappában levő "WINCMD.KEY" fájlt a telepítési pappába másolod.
		<br>Indíthatod a programot.
		<br>
	</div>
	<div class="torrent_leiras">
		Akik eddig megköszönték:<br><br>
		<div id="ncoreKoszonetAjax">
			<p align="center">
				<a href="javascript:void(0);" onclick="thanks(3194806);"><img class="g_koszi"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						alt="Megköszönheted a feltöltést"></a>
			</p>
		</div>
	</div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=3194806&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
</div>
`

const bookHtml = `
<div class="fobox_tartalom">
	<div class="torrent_reszletek_cim">George R. R. Martin - Királyok csatája - A Tűz és Jég dala II.</div>
	<div class="torrent_reszletek_konyvjelzo">
		<a style="font-weight:normal;" id="konyvjelzo_ajax" href="javascript:void(0);" onclick="konyvjelzo('1031296');"
			title="Hozzáadás a könyvjelzőkhöz">[könyvjelzőkhöz adás]</a>
		<a style="font-weight:normal;"
			href="torrents.php?action=addnews&amp;id=1031296&amp;getunique=8ffe570cf24f14b749ee44e32587c118"
			title="A torrent ajánlása a hírfolyamon!">[hírfolyamra]</a>
		<a style="font-weight:normal;" href="torrents.php?action=notification&amp;id=1031296">[bejelentés]</a>
	</div>
	<center>
		<table cellspacing="5" border="0">
			<tbody>
				<tr>
					<td> <a class="fancy_groups" rel="g1031296"
							href="https://nc-img.burst.sh/N-jX-93vQTPizB0v"><img
								src="https://nc-img.burst.sh/AL94kRDr2S2i94Zl" border="0"></a></td>
				</tr>
				<tr>
					<td align="center"><em>(183x298)</em></td>
				</tr>
			</tbody>
		</table>
	</center><br>
	<center id="detailsad">
		<div class="banner" data-id="990" data-banner-id="969" data-zone-id="1,2,3,4"><iframe scrolling="no"
				marginwidth="0" noresize="" marginheight="0" vspace="0" hspace="0"
				src="https://rklm1.parom.hu/dyna13/index.php?ref=ncc" width="800" height="100" frameborder="0"></iframe>
		</div>
	</center>
	<!--<script type="text/javascript">
$(document).ready(function(){
	$('#detailsad').html('<iframe id="detailsadb" src="http://yayo.hu/api/banner/800" style="width:800px;border:0;height:116px"></iframe>');
});
</script>-->
	<br>
	<div class="torrent_reszletek">
		<div class="torrent_col1">
			<div class="dt">Típus:</div>
			<div class="dd"><a title="Még több ebből a kategóriából"
					href="torrents.php?csoport_listazas=osszes_konyv">Ebook</a> &gt; <a
					title="Még több ebből a kategóriából" href="torrents.php?tipus=ebook_hun">eBook/HU</a></div>
			<div class="dt">Feltöltve:</div>
			<div class="dd">2012-06-18 16:39:09</div>
			<div class="dt">Feltöltő:</div>
			<div class="dd">
				<span class="feltolto_szin">Anonymous</span>
			</div>
			<div class="dt">Hozzászólás:</div>
			<div class="dd">3 db</div>
		</div>
		<div class="torrent_col2">
			<div class="dt">Seederek:</div>
			<div class="dd"><a onclick="peers('peers', '1031296', 1);" title="Peerlista"
					href="javascript:void(0);">169</a></div>
			<div class="dt">Leecherek:</div>
			<div class="dd"><a onclick="peers('peers', '1031296', 1);" title="Peerlista"
					href="javascript:void(0);">11</a></div>
			<div class="dt">Letöltve:</div>
			<div class="dd">+++++</div>
			<div class="dt">Sebesség:</div>
			<div class="dd">0 B/s (becsült)</div>
			<div class="dt">Méret:</div>
			<div class="dd">10.75 MiB (11274777 bájt)</div>
			<div class="dt">Tech info:</div>
			<div class="dd"><a onclick="getNfo('nfocontent', '1031296', 1);" title="Mutat"
					href="javascript:void(0);">Mutat</a></div>
			<div class="dt">Fájlok:</div>
			<div class="dd"><a onclick="files('filelist', '1031296', 1);" title="Fájlok"
					href="javascript:void(0);">11</a></div>
			<!--<div class="dt">Mafab.hu</div>
	<div class="dd"><a onclick="getMafab('mafab', 'tt', 0, 1);" href="javascript:void(0);">Mutat</a></div>-->
		</div>
	</div>
	<div style="clear:both;"></div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1031296&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
	<div id="filelist"></div>
	<!-- <div id="mafab"></div>-->
	<div id="otherlist"></div>
	<div id="nfocontent"></div>
	<a name="peers"></a>
	<div id="peers"></div>
	<div id="subtitles"></div>
	<div class="torrent_leiras">
		<table width="100%">
			<tbody>
				<tr>
					<td class="inforbar_img" align="center"><img
							src="https://nc-img.burst.sh/covers/EdeYZgjvCDrCLX-x" alt="Borító"></td>
					<td class="inforbar_txt">
						<table>
							<tbody>
								<tr>
									<td style="vertical-align:top;width:110px;">Címkék:</td>
									<td><a style="border-bottom:1px dotted white;"
											href="torrents.php?tags=fantasy">fantasy</a></td>
								</tr>
							</tbody>
						</table>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class="torrent_leiras proba42">
		<span style="text-decoration:underline">Könyv címe</span>: Királyok csatája - A Tűz és Jég dala II.
		<br><span style="text-decoration:underline">Kiadás dátuma</span>: 2008
		<br><span style="text-decoration:underline">Formátum</span>: <span style="font-size:12pt"><span
				class="highlight">prc, mobi, epub</span></span>
		<br>
		<br>
		<br><span class="highlight">A tűz és jég dalának második könyve.
			<br>
			<br>Hat királyi ház száll síkra a trónért folyó küzdelemben. Az utódlásért vívott harc vérfertőzés és
			apagyilkosság, alkímia és merénylet díszletei között zajlik egy olyan földön, ahol a dicsőség árát vérben
			mérik.</span>
		<br>
		<br>
		<br>George R. R. Martin lebilincselő fantasy eposzának, A tűz és jég dalának második könyve, a Királyok csatája
		az uralkodóját vesztett Westeros birodalmában veszi fel a történet fonalát. Hat királyi ház száll síkra a
		trónért folyó küzdelemben. Az utódlásért vívott harc vérfertőzés és apagyilkosság, alkímia és merénylet
		díszletei között zajlik egy olyan földön, ahol a dicsőség árát vérben mérik. Martin a fantasy mestere.
		<br>
		<br>Kimeríthetetlen képzelőerővel megírt históriái középkori mintákból merítenek ihletet, és a láthatóság
		határát súroló plasztikussággal jelenítik meg fiktív világukat. A sodró, szofisztikált történetvezetés, az
		árnyalt karakterábrázolás és a szexualitásnak a fantasy műfajában szokatlanul nyílt kezelése is fontos szerepet
		játszik abban, hogy Martin sikert sikerre halmozó újítóként vonulhat be a műfaj történetébe.
		<br>
		<br>A tűz és jég dala az elmúlt évtized egyik legjelentősebb fantasy sorozata, minden kötete letehetetlen
		olvasmány.
		<br>
		<br><a href="https://none.st/?q=http://www.libri.hu/konyv/kiralyok-csataja.html" target="_blank"
			class="bb-url">Link</a>
	</div>
	<div class="torrent_leiras">
		Akik eddig megköszönték:<br><br>
		<div id="ncoreKoszonetAjax">
			<p align="center">
				<a href="javascript:void(0);" onclick="thanks(1031296);"><img class="g_koszi"
						src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAP///////wAAACH5BAEAAAIALAAAAAAPAA8AAAINlI+py+0Po5y02otnAQA7"
						alt="Megköszönheted a feltöltést"></a>
			</p>
		</div>
	</div>
	<div class="download"><a style="font-size:11px;text-decoration:underline;text-transform:uppercase;"
			href="torrents.php?action=download&amp;id=1031296&amp;key=7b200e994f68aa46bc4b55c9c85d9fc1">Torrent
			letöltése</a></div>
</div>
`
