// Code generated by github.com/reearth/reearth-backend/tools/cmd/embed, DO NOT EDIT.

package builtin

const pluginManifestJSON_ja string = `{"description":"公式プラグイン","extensions":{"cesium":{"description":"右パネルでシーン全体の設定を変更することができます。タイル、大気、ライティングなどの設定が含まれています。","propertySchema":{"atmosphere":{"description":"地球を覆う大気圏の設定ができます。","fields":{"brightness_shift":{"title":"明度"},"enable_lighting":{"description":"太陽光のON/OFFを切り替えることで、昼夜を表現することができます。","title":"太陽光"},"enable_sun":{"description":"宇宙空間に存在する太陽の表示を切り替えます。","title":"太陽"},"fog":{"description":"霧のON/OFFを切り替えます。","title":"霧"},"fog_density":{"description":"霧の濃度を0以上から設定します。","title":"濃度"},"ground_atmosphere":{"description":"地表の大気圏のON/OFFを切り替えます。","title":"地表の大気"},"hue_shift":{"title":"色相"},"sky_atmosphere":{"description":"地球を覆う大気圏のON/OFFを切り替えます。","title":"上空の大気"},"surturation_shift":{"title":"彩度"}},"title":"大気"},"default":{"fields":{"bgcolor":{"description":"宇宙空間が非表示の場合の、背景色を設定します。","title":"背景色"},"camera":{"description":"ページロード後最初に表示するカメラの位置を設定します。","title":"カメラ初期位置"},"ion":{"description":"自身のCesium IonアカウントからAPIキーを発行し、ここに設定します。Cesium Ionのアセット（タイルデータ、3Dデータなど）の使用が可能になるため、設定を推奨します。","title":"Cesium Icon APIアクセストークン"},"skybox":{"description":"宇宙空間の表示を切り替えます。","title":"宇宙の表示"},"terrain":{"description":"有効にすると、標高データが読み込みこまれ、立体的な地形を表現することができます。","title":"地形"}},"title":"シーン"},"googleAnalytics":{"description":"Google Analyticsを有効にすることで、公開ページがどのように閲覧されているかを分析することが可能です。","fields":{"enableGA":{"description":"Google Analyticsを有効にします。","title":"有効"},"trackingCode":{"description":"ここにグーグルアナリティクスのトラッキングIDを貼り付けることで、公開プロジェクトにこのコードが埋め込まれます。","title":"トラッキングID"}},"title":"Google Analytics"},"tiles":{"description":"手持ちのタイルデータを使用し、地球上に表示することができます。","fields":{"tile_maxLevel":{"title":"最大レベル"},"tile_minLevel":{"title":"最小レベル"},"tile_title":{"title":"名前"},"tile_type":{"choices":{"black_marble":"Black Marble","default":"デフォルト","default_label":"ラベル付き地図","default_road":"道路地図","esri_world_topo":"ESRI Topography","japan_gsi_standard":"地理院地図 標準地図","open_street_map":"Open Street Map","stamen_toner":"Stamen Toner","stamen_watercolor":"Stamen Watercolor","url":"URL"},"title":"種類"},"tile_url":{"title":"URL"}},"title":"タイル"}},"title":"Cesium"},"dlblock":{"description":"表ブロック","propertySchema":{"default":{"fields":{"title":{"title":"タイトル"},"typography":{"title":"フォント"}},"title":"表ブロック"},"items":{"fields":{"item_datanum":{"title":"データ(数字)"},"item_datastr":{"title":"データ(文字)"},"item_datatype":{"choices":{"number":"数字","string":"文字"},"title":"種類"},"item_title":{"title":"タイトル"}},"title":"アイテム"}},"title":"表"},"ellipsoid":{"description":"楕円形ツールを地図上にドラッグ\u0026ドロップすることで追加できます。楕円形ツールによって立体的なオブジェクトを地図上に表示できます。","propertySchema":{"default":{"fields":{"fillColor":{"title":"塗り色"},"height":{"title":"高度"},"position":{"title":"位置"},"radius":{"title":"半径"}},"title":"球体ツール"}},"title":"球体ツール"},"imageblock":{"description":"画像ブロック","propertySchema":{"default":{"fields":{"fullSize":{"title":"フルサイズ"},"image":{"title":"画像"},"imagePositionX":{"choices":{"center":"中央","left":"左","right":"右"},"title":"水平位置"},"imagePositionY":{"choices":{"bottom":"下","center":"中央","top":"上"},"title":"垂直位置"},"imageSize":{"choices":{"contain":"含む","cover":"カバー"},"title":"画像サイズ"},"title":{"title":"タイトル"}},"title":"画像ブロック"}},"title":"画像"},"infobox":{"description":"閲覧者が地図上のレイヤーをクリックした時に表示されるボックスです。テキストや画像、動画などのコンテンツを表示することができます。","propertySchema":{"default":{"fields":{"bgcolor":{"title":"背景色"},"size":{"choices":{"large":"大","small":"小"},"title":"サイズ"},"title":{"title":"タイトル"},"typography":{"title":"フォント"}},"title":"インフォボックス"}},"title":"インフォボックス"},"locationblock":{"description":"位置情報ブロック","propertySchema":{"default":{"fields":{"fullSize":{"title":"フルサイズ"},"location":{"title":"位置情報"},"title":{"title":"タイトル"}},"title":"位置情報ブロック"}},"title":"位置情報"},"marker":{"description":"ドラッグ\u0026ドロップすることで、地図上にマーカーを追加します。マーカーにはテキストや画像を紐づけることができ、閲覧者はマーカーをクリックすることでそれらのコンテンツを見ることができます。","propertySchema":{"default":{"fields":{"extrude":{"title":"地面から線を伸ばす"},"height":{"title":"高度"},"image":{"title":"画像URL"},"imageCrop":{"choices":{"circle":"円形","none":"なし"},"title":"切り抜き"},"imageShadow":{"title":"シャドウ"},"imageShadowBlur":{"title":"シャドウ半径"},"imageShadowColor":{"title":"シャドウ色"},"imageShadowPositionX":{"title":"シャドウX"},"imageShadowPositionY":{"title":"シャドウY"},"imageSize":{"title":"画像サイズ"},"label":{"title":"ラベル"},"labelText":{"title":"ラベル文字"},"labelTypography":{"title":"ラベルフォント"},"location":{"title":"位置"},"pointColor":{"title":"ポイント色"},"pointSize":{"title":"ポイントサイズ"},"style":{"choices":{"image":"アイコン","point":"ポイント"},"title":"表示方法"}},"title":"マーカー"}},"title":"マーカー"},"menu":{"description":"シーンにボタンを設置し、メニューを表示します。追加したボタンに設定されたアクションタイプによって動作が変わります。\\n・リンク：ボタン自体が外部サイトへのリンクになります。\\n・メニュー：追加したメニューを開きます\\n・カメラアクション：クリック時にカメラを移動します。","propertySchema":{"buttons":{"fields":{"buttonBgcolor":{"title":"背景色"},"buttonCamera":{"title":"カメラ"},"buttonColor":{"title":"テキスト色"},"buttonIcon":{"title":"アイコン"},"buttonInvisible":{"title":"非表示"},"buttonLink":{"title":"リンク"},"buttonPosition":{"choices":{"bottomleft":"下左","bottomright":"下右","topleft":"上左","topright":"上右"},"title":"表示位置"},"buttonStyle":{"choices":{"icon":"アイコンのみ","text":"テキストのみ","texticon":"テキスト＋アイコン"},"title":"表示方法"},"buttonTitle":{"title":"タイトル"},"buttonType":{"choices":{"camera":"カメラ移動","link":"リンク","menu":"メニュー開閉"},"title":"アクション"}},"title":"ボタン"},"menu":{"fields":{"menuCamera":{"title":"カメラ"},"menuIcon":{"title":"アイコン"},"menuLink":{"title":"リンク"},"menuTitle":{"title":"タイトル"},"menuType":{"choices":{"border":"区切り線","camera":"カメラ移動","link":"リンク"},"title":"アクション"}},"title":"メニュー"}},"title":"メニュー"},"photooverlay":{"description":"地図上に追加されたフォトオーバーレイを選択すると、設定した画像をモーダル形式で表示することができます。","propertySchema":{"default":{"fields":{"camera":{"description":"クリックされたときに移動するカメラの設定をします。","title":"カメラ"},"height":{"title":"高度"},"image":{"title":"アイコン"},"imageCrop":{"choices":{"circle":"円形","none":"なし"},"title":"切り抜き"},"imageShadow":{"title":"アイコンシャドウ"},"imageShadowBlur":{"title":"シャドウ半径"},"imageShadowColor":{"title":"シャドウ色"},"imageShadowPositionX":{"title":"シャドウX"},"imageShadowPositionY":{"title":"シャドウY"},"imageSize":{"title":"アイコンサイズ"},"location":{"title":"位置"},"photoOverlayImage":{"title":"オーバレイ画像"}},"title":"フォトオーバーレイ"}},"title":"フォトオーバーレイ"},"polygon":{"description":"Polygon primitive","propertySchema":{"default":{"fields":{"fill":{"title":"塗り"},"fillColor":{"title":"塗り色"},"polygon":{"title":"ポリゴン"},"stroke":{"title":"線"},"strokeColor":{"title":"線色"},"strokeWidth":{"title":"線幅"}},"title":"ポリゴン"}},"title":"ポリゴン"},"polyline":{"description":"Polyline primitive","propertySchema":{"default":{"fields":{"coordinates":{"title":"頂点"},"strokeColor":{"title":"線色"},"strokeWidth":{"title":"線幅"}},"title":"直線"}},"title":"直線"},"rect":{"description":"Rectangle primitive","propertySchema":{"default":{"fields":{"extrudedHeight":{"title":"高さ"},"fillColor":{"title":"塗り色"},"height":{"title":"高度"},"image":{"title":"画像URL"},"rect":{"title":"長方形"},"style":{"choices":{"color":"色","image":"画像"},"title":"スタイル"}},"title":"長方形"}},"title":"長方形"},"resource":{"description":"外部からデータ（形式何？？？）をインポートすることができます。地図上に追加後、URLを指定することで外部データが読み込まれます。","propertySchema":{"default":{"fields":{"url":{"choices":{"auto":"自動","czml":"CZML","geojson":"GeoJSON / TopoJSON","kml":"KML"},"title":"ファイル URL"}},"title":"ファイル"}},"title":"ファイル"},"splashscreen":{"description":"ページロード後、最初に表示される演出を設定できます。例えば、プロジェクトのタイトルを閲覧者に見せたり、カメラを移動させることができます。","propertySchema":{"camera":{"fields":{"cameraDelay":{"title":"カメラ移動時間"},"cameraDuration":{"title":"カメラ開始時間"},"cameraPosition":{"title":"カメラ位置"}},"title":"カメラアニメーション"},"overlay":{"fields":{"overlayBgcolor":{"title":"背景色"},"overlayDelay":{"title":"開始時間"},"overlayDuration":{"title":"表示時間"},"overlayEnabled":{"title":"有効"},"overlayImage":{"title":"オーバーレイ画像"},"overlayImageH":{"title":"画像高さ"},"overlayImageW":{"title":"画像幅"},"overlayTransitionDuration":{"title":"フェード時間"}},"title":"オーバーレイ"}},"title":"スプラッシュスクリーン"},"storytelling":{"description":"ストーリーテリング機能を使えば、データ間の繋がりや時系列をもとに、順番に資料を閲覧してもらうことが可能です。使用するには、右パネルから地球上のレイヤーに順番を付与します。","propertySchema":{"default":{"fields":{"autoStart":{"title":"自動再生"},"camera":{"title":"カメラ"},"duration":{"title":"カメラ移動時間"},"range":{"title":"画角"}},"title":"デフォルト"},"stories":{"fields":{"layer":{"title":"レイヤー"},"layerCamera":{"title":"カメラ"},"layerDuration":{"title":"移動時間"},"layerRange":{"title":"カメラ画角"}},"title":"ストーリー"}},"title":"ストーリーテリング"},"textblock":{"description":"Text block","propertySchema":{"default":{"fields":{"markdown":{"title":"マークダウン"},"text":{"title":"コンテンツ"},"title":{"title":"タイトル"},"typography":{"title":"フォント"}},"title":"テキストブロック"}},"title":"テキスト"},"videoblock":{"description":"動画ブロック","propertySchema":{"default":{"fields":{"fullSize":{"title":"フルサイズ"},"title":{"title":"タイトル"},"url":{"title":"動画 URL"}},"title":"動画ブロック"}},"title":"動画"}},"title":"Re:Earth公式プラグイン"}`
