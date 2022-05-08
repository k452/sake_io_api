package persistence

import (
	"reflect"
	"sake_io_api/domain/model"
	"sake_io_api/infra/persistence"
	"sake_io_api/utils"
	"sort"
	"testing"
)

func TestGetRecipeAll(t *testing.T) {
	t.Setenv("GO_ENV", "test")
	up := persistence.NewRecipePersistence(false)
	got := up.GetRecipeAll()
	want1 := []model.Recipe{
		{
			1,
			1,
			"グレンリベットのトゥワイスアップ",
			"「ザ」を冠するホンモノの証。スコッチウイスキーの聖地、スペイサイド地方を代表するシングルモルトウイスキー、それがザ・グレンリベットです。グレンフィディックに継ぎ、世界で2番目に売れているシングルモルト・ウイスキーで、「すべてのシングルモルトはここから始まった」という謳い文句でも知られています。ザ・グレンリベット蒸溜所はフランスのペルノ・リカール社が所有しており、熟成された原酒はシングルモルトの「ザ・グレンリベット」として出荷されるほか、ペルノ・リカール社が生産・販売するシーバス・リーガルやロイヤル・サリュートなどのブレンデッドウイスキーのキーモルトとして使用されています。このページではザ・グレンリベットの歴史からシリーズの各ラインナップを徹底解説します。「ザ・グレンリベットを少しずつ飲み比べできるウイスキーセット」もご用意しているのでぜひご利用ください。",
			"https://prtimes.jp/i/4201/92/resize/d4201-92-201721-0.jpg",
			utils.StringToTime("2020-05-02T16:01:09Z"),
			utils.StringToTime("2020-05-02T16:01:09Z"),
		},
		{
			2,
			1,
			"クラフト🍺入門",
			"クラフトビールの定義は諸説ありますが、私たちは以下のように考えます。クラフト（craft）は英語で「技術」「工芸」「職人技」などを意味する言葉。クラフトビールとは、小規模な醸造所がつくる多様で個性的なビールを指します。これまでにない多様性と、個性的な味わいやブランドを備えているのが特徴です。日本では1994年の酒税法改正で規制緩和されたことにより、小規模なビール醸造所が日本全国に登場しました。私たち「よなよなの里（企業名：ヤッホーブルーイング）」の醸造所もその中のひとつです。よなよなの里の醸造所のレギュラーラインナップはこの4種類。",
			"https://mycraftbeers.com/wp-content/uploads/2017/05/shutterstock_396628984-322.jpg",
			utils.StringToTime("2020-06-02T16:01:09Z"),
			utils.StringToTime("2020-06-02T16:01:09Z"),
		},
		{
			3,
			3,
			"今呑むべき日本酒",
			"おもちゃメーカー・バンダイの「ガシャポン」と日本酒専門WEBメディア「SAKETIMES」がコラボしたコレクションフィギュア『日本の銘酒 SAKE COLLECTION』が、2021年4月第4週から順次、販売されます。 SAKETIMESが厳選した全国の銘酒の一升瓶をミニチュア化（全長約6㎝）し、ガシャポンとして、全国のおもちゃ売場・量販店・家電店などに設置されたカプセル自販機にて販売します。",
			"https://jp.sake-times.com/wp-content/uploads/2021/04/sake_nihonshu-gacyapon-08.png",
			utils.StringToTime("2021-01-07T16:01:09Z"),
			utils.StringToTime("2021-01-07T16:01:09Z"),
		},
		{
			4,
			2,
			"霧島酒造5選",
			"芋焼酎の世界に革命を起こした霧島酒造。初蔵出しは大正5年だそうです。サツマイモの本場である、地元宮崎で獲れた良質なサツマイモと、霧島連山に降った雨がシラス台地で濾過された美味しい水、霧島裂罅水（きりしまれっかすい）で仕込んだ芋焼酎で一世を風靡しました。 製品の中でも「黒霧（くろきり）」の愛称で親しまれる「黒霧島」は、2000年代初頭の芋焼酎ブームを牽引し、国民的な人気酒になりました。今でも今日は焼酎飲むか、という時、酒屋さんでついつい黒霧を手に取ってしまう方も多いのではないでしょうか。 霧島酒造の芋焼酎には、人気の「黒霧島」以外にも「〇✕霧島」という兄弟のようなお酒が色々あります。ついつい、いつものクセで「黒」を買ってしまいがちですが、本当のところ他の霧島ってどんな味なの？　というわけで霧島シリーズのうち、手軽に町の酒屋さんで手に入る銘柄を飲み比べてみることに。",
			"https://www.kirishima.co.jp/assets/images/ximg_products_bottle.png.pagespeed.ic.pRB9_IBZ5z.webp",
			utils.StringToTime("2021-01-07T16:01:09Z"),
			utils.StringToTime("2021-01-07T16:01:09Z"),
		},
		{
			5,
			4,
			"肝臓治療法",
			"「二日酔い」とは、お酒を飲みすぎた翌日に起こる不快な症状のことです。二日酔いはアルコールが分解されてできる「アセトアルデヒド」と呼ばれる物質が、肝臓で処理されないために起こります。実際に、アルコールは体内でどのように代謝されるのでしょうか？アルコールを飲むと、胃から20％、小腸から80％吸収されます。吸収されたアルコールのほとんどは肝臓へ運ばれ、代謝されていきます。肝臓の中では、酵素の働きでアセトアルデヒドという物質になり、さらにアセテート（酢酸）に分解され、血液にのって全身を回ります。このアセテートは、筋肉や脂肪組織で水と二酸化炭素に分解され、体外へ排出されていきます。お酒をたくさん飲みすぎてしまった場合、肝臓でアセトアルデヒドを十分に処理することができず、血液中のアセトアルデヒドの濃度が高くなってしまいます。実は、このアセトアルデヒドは毒性があり、濃度が高まると吐き気や動悸、頭痛などを引き起こします。これが二日酔いの原因となります。",
			"https://www.hepa-w.jp/img/top/hepa2.png",
			utils.StringToTime("2021-05-23T16:01:09Z"),
			utils.StringToTime("2021-05-23T16:01:09Z"),
		},
	}
	want2 := []model.Recipe{}

	tests := []struct {
		name string
		want []model.Recipe
	}{
		{
			name: "recipeテーブル全データ取得が成功",
			want: want1,
		},
		{
			name: "recipeテーブル全データ取得が失敗",
			want: want2,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i == 0 {
				if !reflect.DeepEqual(sortRecipeAll(got), sortRecipeAll(tt.want)) {
					t.Error("error")
				}
			} else if i == 1 {
				if reflect.DeepEqual(sortRecipeAll(got), sortRecipeAll(tt.want)) {
					t.Error("error")
				}
			}
		})
	}
}

func sortRecipeAll(recipeSlice []model.Recipe) []model.Recipe {
	sort.Slice(recipeSlice, func(i, j int) bool {
		return recipeSlice[i].Recipe_id < recipeSlice[j].Recipe_id
	})
	return recipeSlice
}