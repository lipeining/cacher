package main

type AttrRDB struct {
	Name               map[string]string   `json:"name"`
	AttrId             string              `json:"attr_id"`
	Icon               string              `json:"icon"`
	IconRegion         map[string]string   `json:"icon_region"`
	Color              string              `json:"color"`
	CopyrightOwner     string              `json:"copyright_owner"`
	Creator            string              `json:"creator"`
	EndedAt            int64               `json:"ended_at"`
	StartedAt          int64               `json:"started_at"`
	IsNew              int                 `json:"is_new"`
	IsNewTime          int64               `json:"is_new_time"`
	Sort               int                 `json:"sort"`
	PaidType           int                 `json:"paid_type"`
	Version            []string            `json:"version"`
	RegionType         int                 `json:"region_type"`
	Region             map[string]bool     `json:"region"`
	ReleaseMode        int                 `json:"release_mode"`
	DeleteAfterOffline int                 `json:"delete_after_offline"`
	DownloadType       int                 `json:"download_type"`
	CreateAt           int64               `json:"create_at"`
	MulGId             []string            `json:"mul_g_id"`
	Tags               map[string][]string `json:"tags"`
	VideoEditor        int                 `json:"video_editor"`
	PictureEditor      int                 `json:"picture_editor"`
	PuriPlus           int                 `json:"puri_plus"`
	Delivery           int                 `json:"delivery"`
	CustMetadata       map[string]string   `json:"cust_metadata"`
}

type CategoryRDB struct {
	BgColor            string              `json:"bg_color"`
	BestSort           int                 `json:"best_sort"`
	CategoryId         string              `json:"category_id"`
	Color              string              `json:"color"`
	CoverImg           string              `json:"cover_img"`
	CreateAt           int64               `json:"create_at"`
	CopyrightOwner     string              `json:"copyright_owner"`
	Creator            string              `json:"creator"`
	DisplayImg         string              `json:"display_img"`
	DeleteAfterOffline int                 `json:"delete_after_offline"`
	DownloadType       int                 `json:"download_type"`
	EndedAt            int64               `json:"ended_at"`
	IsNew              int                 `json:"is_new"`
	IsNewTime          int64               `json:"is_new_time"`
	IsListDisplay      int                 `json:"is_list_display"`
	Icon               string              `json:"icon"`
	IsColorSwitch      int                 `json:"is_color_switch"`
	IsHot              int                 `json:"is_hot"`
	IsHotSort          int                 `json:"is_hot_sort"`
	IsListBest         int                 `json:"is_list_best"`
	IntroductoryCopy   map[string]string   `json:"introductory_copy"`
	IsSub              int                 `json:"is_sub"`
	MulGId             []string            `json:"mul_g_id"`
	Name               map[string]string   `json:"name"`
	OldId              string              `json:"old_id"`
	OldSkuId           string              `json:"old_sku_id"`
	PaidType           int                 `json:"paid_type"`
	PaidSort           int                 `json:"paid_sort"`
	RegionType         int                 `json:"region_type"`
	Region             map[string]bool     `json:"region"`
	Scenes             string              `json:"scenes"`
	StartedAt          int64               `json:"started_at"`
	ReleaseMode        int                 `json:"release_mode"`
	Sort               int                 `json:"sort"`
	Sub                []string            `json:"sub"`
	Tags               map[string][]string `json:"tags"`
	Version            []string            `json:"version"`
	UpperIds           []string            `json:"upper_ids"`
	IsListBestTime     int64               `json:"is_list_best_time"`
	IsHotTime          int64               `json:"is_hot_time"`
	IsListDisplayTime  int64               `json:"is_list_display_time"`
	VideoEditor        int                 `json:"video_editor"`
	PictureEditor      int                 `json:"picture_editor"`
	PuriPlus           int                 `json:"puri_plus"`
	Delivery           int                 `json:"delivery"`
	IconRegion         map[string]string   `json:"icon_region"`
	PreviewImgRegion   map[string]string   `json:"preview_img_region"`
	CoverImgRegion     map[string]string   `json:"cover_img_region"`
	ScenesRegion       map[string][]string `json:"scenes_region"`
	DisplayImgRegion   map[string][]string `json:"display_img_region"`
	StickerType        int                 `json:"sticker_type"`
	CustMetadata       map[string]string   `json:"cust_metadata"`
}

type GoodsRDB struct {
	Alpha              map[string]int      `json:"alpha"`
	BgColor            string              `json:"bg_color"`
	BestSort           int                 `json:"best_sort"`
	BgColorModel       int                 `json:"bg_color_model"`
	ClientMaterialPay  int                 `json:"client_material_pay"`
	Color              string              `json:"color"`
	CreateAt           int64               `json:"create_at"`
	ClientShow         int                 `json:"client_show"`
	CoverImgRegion     map[string]string   `json:"cover_img_region"`
	DeleteAfterOffline int                 `json:"delete_after_offline"`
	DownloadType       int                 `json:"download_type"`
	DisplayImgRegion   map[string][]string `json:"display_img_region"`
	EndedAt            int64               `json:"ended_at"`
	GId                string              `json:"g_id"`
	GroupType          int                 `json:"group_type"`
	IsHot              int                 `json:"is_hot"`
	IsHotSort          int                 `json:"is_hot_sort"`
	IsListBest         int                 `json:"is_list_best"`
	IsNew              int                 `json:"is_new"`
	IsNewTime          int64               `json:"is_new_time"`
	IsListDisplay      int                 `json:"is_list_display"`
	Icon               string              `json:"icon"`
	IconRegion         map[string]string   `json:"icon_region"`
	IconRatio          string              `json:"icon_ratio"`
	IconProportion     string              `json:"icon_proportion"`
	IsListBestTime     int64               `json:"is_list_best_time"`
	IsHotTime          int64               `json:"is_hot_time"`
	IsListDisplayTime  int64               `json:"is_list_display_time"`
	MId                string              `json:"m_id"`
	MulAttrId          []string            `json:"mul_attr_id"`
	MaterialIds        []string            `json:"material_ids"`
	MainLayerCount     int                 `json:"main_layer_count"`
	Name               map[string]string   `json:"name"`
	OverlayMode        int                 `json:"overlay_mode"`
	OldId              string              `json:"old_id"`
	PreviewImg         string              `json:"preview_img"`
	OriginalImg        map[string][]string `json:"original_img"`
	PaidType           int                 `json:"paid_type"`
	PaidSort           int                 `json:"paid_sort"`
	PresetValue        []GPresetValue       `json:"preset_value"`
	PreviewImgRegion   map[string]string   `json:"preview_img_region"`
	RelationIcons      []map[string]string `json:"relation_icons"`
	ReleaseMode        int                 `json:"release_mode"`
	RegionType         int                 `json:"region_type"`
	Region             map[string]bool     `json:"region"`
	StartedAt          int64               `json:"started_at"`
	Scenes             string              `json:"scenes"`
	Source             int                 `json:"source"`
	Sort               int                 `json:"sort"`
	Sub                []GProduct           `json:"sub"`
	ScenesRegion       map[string][]string `json:"scenes_region"`
	Tags               map[string][]string `json:"tags"`
	TbgType            int                 `json:"tbg_type"`
	UpperIds           []string            `json:"upper_ids"`
	Version            []string            `json:"version"`
	PictureEditor      int                 `json:"picture_editor"`
	VideoEditor        int                 `json:"video_editor"`
	Delivery           int                 `json:"delivery"`
	PuriPlus           int                 `json:"puri_plus"`
	CustMetadata       map[string]string   `json:"cust_metadata"`
	//素材
	AssetVersion       string  `json:"asset_version"`
	AnimationType      int     `json:"animation_type"`
	CoverVideo         string  `json:"cover_video"`
	CopyrightOwner     string  `json:"copyright_owner"`
	Creator            string  `json:"creator"`
	Duration           int64   `json:"duration"`
	Dependent          int     `json:"dependent"`
	Effect             GEffect  `json:"effect"`
	File               GFileRDB `json:"file"`
	ImgRatio           string  `json:"img_ratio"`
	IsColorSwitch      int     `json:"is_color_switch"`
	IsGl3              int     `json:"is_gl3"`
	IsTouch            int     `json:"is_touch"`
	IsMask             int     `json:"is_mask"`
	IsPortrait         int     `json:"is_portrait"`
	PreviewVideo       string  `json:"preview_video"`
	Sex                int     `json:"sex"`
	StickerType        int     `json:"sticker_type"`
	StartPoint         int64   `json:"start_point"`
	Singer             string  `json:"singer"`
	TransitionPosition int     `json:"transition_position"`
	UseScenes          int     `json:"use_scenes"`
	Prompt             string  `json:"prompt"`
	Point              int     `json:"point"`
	ApplicationId      int     `json:"application_id"`
}

type GProduct struct {
	MId           string            `json:"m_id"`
	OldId         string            `json:"old_id"`
	AssetId       string            `json:"asset_id"`
	AssetUrl      string            `json:"asset_url"`
	Type          int               `json:"type"`
	PaidType      int               `json:"paid_type"`
	ApplicationId int               `json:"application_id"`
	Name          map[string]string `json:"name"`
	Icon          map[string]string `json:"icon"`
}

type GFileRDB struct {
	FileSize string `json:"file_size"`
	FileUrl  string `json:"file_url"`
	FileId   string `json:"file_id"`
	MId      string `json:"m_id"`
}

type GPresetValue struct {
	Key          string            `json:"key"`
	Name         map[string]string `json:"name"`
	Degree       map[string]int    `json:"degree"`
	Min          int               `json:"min"`
	Max          int               `json:"max"`
	Type         int               `json:"type"`
	DefaultValue int               `json:"default_value"`
	Status       int               `json:"status"`
}

type GEffect struct {
	Type          int    `json:"type"`
	Online        string `json:"online"`
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	EffectAddress string `json:"effect_address"`
	StyleId       string `json:"style_id"`
}
