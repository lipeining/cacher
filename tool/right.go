package main

type AttrInfo struct {
	Name               map[string]string   `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AttrId             string              `protobuf:"bytes,2,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	Icon               string              `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	IconRegion         map[string]string   `protobuf:"bytes,4,rep,name=icon_region,json=iconRegion,proto3" json:"icon_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Color              string              `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	CopyrightOwner     string              `protobuf:"bytes,6,opt,name=copyright_owner,json=copyrightOwner,proto3" json:"copyright_owner,omitempty"`
	Creator            string              `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	EndedAt            int64               `protobuf:"varint,8,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	StartedAt          int64               `protobuf:"varint,9,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	IsNew              int32               `protobuf:"varint,10,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsNewTime          int64               `protobuf:"varint,11,opt,name=is_new_time,json=isNewTime,proto3" json:"is_new_time,omitempty"`
	Sort               int32               `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`
	PaidType           int32               `protobuf:"varint,13,opt,name=paid_type,json=paidType,proto3" json:"paid_type,omitempty"`
	Version            []string            `protobuf:"bytes,14,rep,name=version,proto3" json:"version,omitempty"`
	RegionType         int32               `protobuf:"varint,15,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	ReleaseMode        int32               `protobuf:"varint,16,opt,name=release_mode,json=releaseMode,proto3" json:"release_mode,omitempty"`
	DeleteAfterOffline int32               `protobuf:"varint,17,opt,name=delete_after_offline,json=deleteAfterOffline,proto3" json:"delete_after_offline,omitempty"`
	DownloadType       int32               `protobuf:"varint,18,opt,name=download_type,json=downloadType,proto3" json:"download_type,omitempty"`
	CreateAt           int64               `protobuf:"varint,19,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	MulGId             []string            `protobuf:"bytes,20,rep,name=mul_g_id,json=mulGId,proto3" json:"mul_g_id,omitempty"`
	Tags               map[string][]string `protobuf:"bytes,21,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VideoEditor        int32               `protobuf:"varint,22,opt,name=video_editor,json=videoEditor,proto3" json:"video_editor,omitempty"`
	PictureEditor      int32               `protobuf:"varint,23,opt,name=picture_editor,json=pictureEditor,proto3" json:"picture_editor,omitempty"`
	PuriPlus           int32               `protobuf:"varint,24,opt,name=puri_plus,json=puriPlus,proto3" json:"puri_plus,omitempty"`
	Delivery           int32               `protobuf:"varint,25,opt,name=delivery,proto3" json:"delivery,omitempty"`
	CustMetadata       map[string]string   `protobuf:"bytes,26,rep,name=cust_metadata,json=custMetadata,proto3" json:"cust_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type CategoryInfo struct {
	BgColor            string              `protobuf:"bytes,1,opt,name=bg_color,json=bgColor,proto3" json:"bg_color,omitempty"`
	BestSort           int32               `protobuf:"varint,2,opt,name=best_sort,json=bestSort,proto3" json:"best_sort,omitempty"`
	CategoryId         string              `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Color              string              `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	CoverImg           string              `protobuf:"bytes,5,opt,name=cover_img,json=coverImg,proto3" json:"cover_img,omitempty"`
	CreateAt           int64               `protobuf:"varint,6,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	CopyrightOwner     string              `protobuf:"bytes,7,opt,name=copyright_owner,json=copyrightOwner,proto3" json:"copyright_owner,omitempty"`
	Creator            string              `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	DisplayImg         string              `protobuf:"bytes,9,opt,name=display_img,json=displayImg,proto3" json:"display_img,omitempty"`
	DeleteAfterOffline int32               `protobuf:"varint,10,opt,name=delete_after_offline,json=deleteAfterOffline,proto3" json:"delete_after_offline,omitempty"`
	DownloadType       int32               `protobuf:"varint,11,opt,name=download_type,json=downloadType,proto3" json:"download_type,omitempty"`
	EndedAt            int64               `protobuf:"varint,12,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	IsNew              int32               `protobuf:"varint,13,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsNewTime          int64               `protobuf:"varint,14,opt,name=is_new_time,json=isNewTime,proto3" json:"is_new_time,omitempty"`
	IsListDisplay      int32               `protobuf:"varint,15,opt,name=is_list_display,json=isListDisplay,proto3" json:"is_list_display,omitempty"`
	Icon               string              `protobuf:"bytes,16,opt,name=icon,proto3" json:"icon,omitempty"`
	IsColorSwitch      int32               `protobuf:"varint,17,opt,name=is_color_switch,json=isColorSwitch,proto3" json:"is_color_switch,omitempty"`
	IsHot              int32               `protobuf:"varint,18,opt,name=is_hot,json=isHot,proto3" json:"is_hot,omitempty"`
	IsHotSort          int32               `protobuf:"varint,19,opt,name=is_hot_sort,json=isHotSort,proto3" json:"is_hot_sort,omitempty"`
	IsListBest         int32               `protobuf:"varint,20,opt,name=is_list_best,json=isListBest,proto3" json:"is_list_best,omitempty"`
	IntroductoryCopy   map[string]string   `protobuf:"bytes,21,rep,name=introductory_copy,json=introductoryCopy,proto3" json:"introductory_copy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsSub              int32               `protobuf:"varint,22,opt,name=is_sub,json=isSub,proto3" json:"is_sub,omitempty"`
	MulGId             []string            `protobuf:"bytes,23,rep,name=mul_g_id,json=mulGId,proto3" json:"mul_g_id,omitempty"`
	Name               map[string]string   `protobuf:"bytes,24,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OldId              string              `protobuf:"bytes,25,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	OldSkuId           string              `protobuf:"bytes,26,opt,name=old_sku_id,json=oldSkuId,proto3" json:"old_sku_id,omitempty"`
	PaidType           int32               `protobuf:"varint,27,opt,name=paid_type,json=paidType,proto3" json:"paid_type,omitempty"`
	PaidSort           int32               `protobuf:"varint,28,opt,name=paid_sort,json=paidSort,proto3" json:"paid_sort,omitempty"`
	RegionType         int32               `protobuf:"varint,29,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	Region             map[string]bool     `protobuf:"bytes,30,rep,name=region,proto3" json:"region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Scenes             string              `protobuf:"bytes,31,opt,name=scenes,proto3" json:"scenes,omitempty"`
	StartedAt          int64               `protobuf:"varint,32,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ReleaseMode        int32               `protobuf:"varint,33,opt,name=release_mode,json=releaseMode,proto3" json:"release_mode,omitempty"`
	Sort               int32               `protobuf:"varint,34,opt,name=sort,proto3" json:"sort,omitempty"`
	Sub                []string            `protobuf:"bytes,35,rep,name=sub,proto3" json:"sub,omitempty"`
	Tags               map[string][]string `protobuf:"bytes,36,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Version            []string            `protobuf:"bytes,37,rep,name=version,proto3" json:"version,omitempty"`
	UpperIds           []string            `protobuf:"bytes,38,rep,name=upper_ids,json=upperIds,proto3" json:"upper_ids,omitempty"`
	IsListBestTime     int64               `protobuf:"varint,39,opt,name=is_list_best_time,json=isListBestTime,proto3" json:"is_list_best_time,omitempty"`
	IsHotTime          int64               `protobuf:"varint,40,opt,name=is_hot_time,json=isHotTime,proto3" json:"is_hot_time,omitempty"`
	IsListDisplayTime  int64               `protobuf:"varint,41,opt,name=is_list_display_time,json=isListDisplayTime,proto3" json:"is_list_display_time,omitempty"`
	VideoEditor        int32               `protobuf:"varint,42,opt,name=video_editor,json=videoEditor,proto3" json:"video_editor,omitempty"`
	PictureEditor      int32               `protobuf:"varint,43,opt,name=picture_editor,json=pictureEditor,proto3" json:"picture_editor,omitempty"`
	PuriPlus           int32               `protobuf:"varint,44,opt,name=puri_plus,json=puriPlus,proto3" json:"puri_plus,omitempty"`
	Delivery           int32               `protobuf:"varint,45,opt,name=delivery,proto3" json:"delivery,omitempty"`
	IconRegion         map[string]string   `protobuf:"bytes,46,rep,name=icon_region,json=iconRegion,proto3" json:"icon_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PreviewImgRegion   map[string]string   `protobuf:"bytes,47,rep,name=preview_img_region,json=previewImgRegion,proto3" json:"preview_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CoverImgRegion     map[string]string   `protobuf:"bytes,48,rep,name=cover_img_region,json=coverImgRegion,proto3" json:"cover_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ScenesRegion       map[string][]string `protobuf:"bytes,49,rep,name=scenes_region,json=scenesRegion,proto3" json:"scenes_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DisplayImgRegion   map[string][]string `protobuf:"bytes,50,rep,name=display_img_region,json=displayImgRegion,proto3" json:"display_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StickerType        int32               `protobuf:"varint,51,opt,name=sticker_type,json=stickerType,proto3" json:"sticker_type,omitempty"`
	CustMetadata       map[string]string   `protobuf:"bytes,52,rep,name=cust_metadata,json=custMetadata,proto3" json:"cust_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type GoodsInfo struct {
	Alpha              map[string]int32    `protobuf:"bytes,1,rep,name=alpha,proto3" json:"alpha,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	BgColor            string              `protobuf:"bytes,2,opt,name=bg_color,json=bgColor,proto3" json:"bg_color,omitempty"`
	BestSort           int32               `protobuf:"varint,3,opt,name=best_sort,json=bestSort,proto3" json:"best_sort,omitempty"`
	BgColorModel       int32               `protobuf:"varint,4,opt,name=bg_color_model,json=bgColorModel,proto3" json:"bg_color_model,omitempty"`
	ClientMaterialPay  int32               `protobuf:"varint,5,opt,name=client_material_pay,json=clientMaterialPay,proto3" json:"client_material_pay,omitempty"`
	Color              string              `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	CreateAt           int64               `protobuf:"varint,7,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	ClientShow         int32               `protobuf:"varint,8,opt,name=client_show,json=clientShow,proto3" json:"client_show,omitempty"`
	CoverImgRegion     map[string]string   `protobuf:"bytes,9,rep,name=cover_img_region,json=coverImgRegion,proto3" json:"cover_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DeleteAfterOffline int32               `protobuf:"varint,10,opt,name=delete_after_offline,json=deleteAfterOffline,proto3" json:"delete_after_offline,omitempty"`
	DownloadType       int32               `protobuf:"varint,11,opt,name=download_type,json=downloadType,proto3" json:"download_type,omitempty"`
	DisplayImgRegion   map[string][]string `protobuf:"bytes,12,rep,name=display_img_region,json=displayImgRegion,proto3" json:"display_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EndedAt            int64               `protobuf:"varint,13,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	GId                string              `protobuf:"bytes,14,opt,name=g_id,json=gId,proto3" json:"g_id,omitempty"`
	GroupType          int32               `protobuf:"varint,15,opt,name=group_type,json=groupType,proto3" json:"group_type,omitempty"`
	IsHot              int32               `protobuf:"varint,16,opt,name=is_hot,json=isHot,proto3" json:"is_hot,omitempty"`
	IsHotSort          int32               `protobuf:"varint,17,opt,name=is_hot_sort,json=isHotSort,proto3" json:"is_hot_sort,omitempty"`
	IsListBest         int32               `protobuf:"varint,18,opt,name=is_list_best,json=isListBest,proto3" json:"is_list_best,omitempty"`
	IsNew              int32               `protobuf:"varint,19,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	IsNewTime          int64               `protobuf:"varint,20,opt,name=is_new_time,json=isNewTime,proto3" json:"is_new_time,omitempty"`
	IsListDisplay      int32               `protobuf:"varint,21,opt,name=is_list_display,json=isListDisplay,proto3" json:"is_list_display,omitempty"`
	Icon               string              `protobuf:"bytes,22,opt,name=icon,proto3" json:"icon,omitempty"`
	IconRegion         map[string]string   `protobuf:"bytes,23,rep,name=icon_region,json=iconRegion,proto3" json:"icon_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IconRatio          string              `protobuf:"bytes,24,opt,name=icon_ratio,json=iconRatio,proto3" json:"icon_ratio,omitempty"`
	IconProportion     string              `protobuf:"bytes,25,opt,name=icon_proportion,json=iconProportion,proto3" json:"icon_proportion,omitempty"`
	IsListBestTime     int32               `protobuf:"varint,26,opt,name=is_list_best_time,json=isListBestTime,proto3" json:"is_list_best_time,omitempty"`
	IsHotTime          int32               `protobuf:"varint,27,opt,name=is_hot_time,json=isHotTime,proto3" json:"is_hot_time,omitempty"`
	IsListDisplayTime  int32               `protobuf:"varint,28,opt,name=is_list_display_time,json=isListDisplayTime,proto3" json:"is_list_display_time,omitempty"`
	MId                string              `protobuf:"bytes,29,opt,name=m_id,json=mId,proto3" json:"m_id,omitempty"`
	MulAttrId          []string            `protobuf:"bytes,30,rep,name=mul_attr_id,json=mulAttrId,proto3" json:"mul_attr_id,omitempty"`
	MaterialIds        []string            `protobuf:"bytes,31,rep,name=material_ids,json=materialIds,proto3" json:"material_ids,omitempty"`
	MainLayerCount     int32               `protobuf:"varint,32,opt,name=main_layer_count,json=mainLayerCount,proto3" json:"main_layer_count,omitempty"`
	Name               map[string]string   `protobuf:"bytes,33,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OverlayMode        int32               `protobuf:"varint,34,opt,name=overlay_mode,json=overlayMode,proto3" json:"overlay_mode,omitempty"`
	OldId              string              `protobuf:"bytes,35,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	PreviewImg         string              `protobuf:"bytes,36,opt,name=preview_img,json=previewImg,proto3" json:"preview_img,omitempty"`
	OriginalImg        map[string][]string `protobuf:"bytes,37,rep,name=original_img,json=originalImg,proto3" json:"original_img,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PaidType           int32               `protobuf:"varint,38,opt,name=paid_type,json=paidType,proto3" json:"paid_type,omitempty"`
	PaidSort           int32               `protobuf:"varint,39,opt,name=paid_sort,json=paidSort,proto3" json:"paid_sort,omitempty"`
	PresetValue        []*GoodsPresetValue `protobuf:"bytes,40,rep,name=preset_value,json=presetValue,proto3" json:"preset_value,omitempty"`
	PreviewImgRegion   map[string]string   `protobuf:"bytes,41,rep,name=preview_img_region,json=previewImgRegion,proto3" json:"preview_img_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RelationIcons      []map[string]string `protobuf:"bytes,42,rep,name=relation_icons,json=relationIcons,proto3" json:"relation_icons,omitempty"`
	ReleaseMode        int32               `protobuf:"varint,43,opt,name=release_mode,json=releaseMode,proto3" json:"release_mode,omitempty"`
	RegionType         int32               `protobuf:"varint,44,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	Region             map[string]bool     `protobuf:"bytes,45,rep,name=region,proto3" json:"region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	StartedAt          int64               `protobuf:"varint,46,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Scenes             string              `protobuf:"bytes,47,opt,name=scenes,proto3" json:"scenes,omitempty"`
	Source             int32               `protobuf:"varint,48,opt,name=source,proto3" json:"source,omitempty"`
	Sort               int32               `protobuf:"varint,49,opt,name=sort,proto3" json:"sort,omitempty"`
	Sub                []*GoodsProduct     `protobuf:"bytes,50,rep,name=sub,proto3" json:"sub,omitempty"`
	ScenesRegion       map[string][]string `protobuf:"bytes,51,rep,name=scenes_region,json=scenesRegion,proto3" json:"scenes_region,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags               map[string][]string `protobuf:"bytes,52,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TbgType            int32               `protobuf:"varint,53,opt,name=tbg_type,json=tbgType,proto3" json:"tbg_type,omitempty"`
	UpperIds           []string            `protobuf:"bytes,54,rep,name=upper_ids,json=upperIds,proto3" json:"upper_ids,omitempty"`
	Version            []string            `protobuf:"bytes,55,rep,name=version,proto3" json:"version,omitempty"`
	PictureEditor      int32               `protobuf:"varint,56,opt,name=picture_editor,json=pictureEditor,proto3" json:"picture_editor,omitempty"`
	VideoEditor        int32               `protobuf:"varint,57,opt,name=video_editor,json=videoEditor,proto3" json:"video_editor,omitempty"`
	Delivery           int32               `protobuf:"varint,58,opt,name=delivery,proto3" json:"delivery,omitempty"`
	PuriPlus           int32               `protobuf:"varint,59,opt,name=puri_plus,json=puriPlus,proto3" json:"puri_plus,omitempty"`
	CustMetadata       map[string]string   `protobuf:"bytes,60,rep,name=cust_metadata,json=custMetadata,proto3" json:"cust_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AssetVersion       string              `protobuf:"bytes,61,opt,name=asset_version,json=assetVersion,proto3" json:"asset_version,omitempty"`
	AnimationType      int32               `protobuf:"varint,62,opt,name=animation_type,json=animationType,proto3" json:"animation_type,omitempty"`
	CoverVideo         string              `protobuf:"bytes,63,opt,name=cover_video,json=coverVideo,proto3" json:"cover_video,omitempty"`
	CopyrightOwner     string              `protobuf:"bytes,64,opt,name=copyright_owner,json=copyrightOwner,proto3" json:"copyright_owner,omitempty"`
	Creator            string              `protobuf:"bytes,65,opt,name=creator,proto3" json:"creator,omitempty"`
	Duration           int64               `protobuf:"varint,66,opt,name=duration,proto3" json:"duration,omitempty"`
	Dependent          int32               `protobuf:"varint,67,opt,name=dependent,proto3" json:"dependent,omitempty"`
	Effect             *GoodsEffect        `protobuf:"bytes,68,opt,name=effect,proto3" json:"effect,omitempty"`
	File               *GoodsFileRDB       `protobuf:"bytes,69,opt,name=file,proto3" json:"file,omitempty"`
	ImgRatio           string              `protobuf:"bytes,70,opt,name=img_ratio,json=imgRatio,proto3" json:"img_ratio,omitempty"`
	IsColorSwitch      int32               `protobuf:"varint,71,opt,name=is_color_switch,json=isColorSwitch,proto3" json:"is_color_switch,omitempty"`
	IsGl3              int32               `protobuf:"varint,72,opt,name=is_gl3,json=isGl3,proto3" json:"is_gl3,omitempty"`
	IsTouch            int32               `protobuf:"varint,73,opt,name=is_touch,json=isTouch,proto3" json:"is_touch,omitempty"`
	IsMask             int32               `protobuf:"varint,74,opt,name=is_mask,json=isMask,proto3" json:"is_mask,omitempty"`
	IsPortrait         int32               `protobuf:"varint,75,opt,name=is_portrait,json=isPortrait,proto3" json:"is_portrait,omitempty"`
	PreviewVideo       string              `protobuf:"bytes,76,opt,name=preview_video,json=previewVideo,proto3" json:"preview_video,omitempty"`
	Sex                int32               `protobuf:"varint,77,opt,name=sex,proto3" json:"sex,omitempty"`
	StickerType        int32               `protobuf:"varint,78,opt,name=sticker_type,json=stickerType,proto3" json:"sticker_type,omitempty"`
	StartPoint         int64               `protobuf:"varint,79,opt,name=start_point,json=startPoint,proto3" json:"start_point,omitempty"`
	Singer             string              `protobuf:"bytes,80,opt,name=singer,proto3" json:"singer,omitempty"`
	TransitionPosition int32               `protobuf:"varint,81,opt,name=transition_position,json=transitionPosition,proto3" json:"transition_position,omitempty"`
	UseScenes          int32               `protobuf:"varint,82,opt,name=use_scenes,json=useScenes,proto3" json:"use_scenes,omitempty"`
	Prompt             string              `protobuf:"bytes,83,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Point              int32               `protobuf:"varint,84,opt,name=point,proto3" json:"point,omitempty"`
	ApplicationId      int32               `protobuf:"varint,85,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

type GoodsPresetValue struct {
	Key          string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name         map[string]string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Degree       map[string]int32  `protobuf:"bytes,3,rep,name=degree,proto3" json:"degree,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Min          int32             `protobuf:"varint,4,opt,name=min,proto3" json:"min,omitempty"`
	Max          int32             `protobuf:"varint,5,opt,name=max,proto3" json:"max,omitempty"`
	Type         int32             `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	DefaultValue int32             `protobuf:"varint,7,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	Status       int32             `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
}

type GoodsProduct struct {
	MId           string            `protobuf:"bytes,1,opt,name=m_id,json=mId,proto3" json:"m_id,omitempty"`
	OldId         string            `protobuf:"bytes,2,opt,name=old_id,json=oldId,proto3" json:"old_id,omitempty"`
	AssetId       string            `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssetUrl      string            `protobuf:"bytes,4,opt,name=asset_url,json=assetUrl,proto3" json:"asset_url,omitempty"`
	Type          int32             `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	PaidType      int32             `protobuf:"varint,6,opt,name=paid_type,json=paidType,proto3" json:"paid_type,omitempty"`
	ApplicationId int32             `protobuf:"varint,7,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Name          map[string]string `protobuf:"bytes,8,rep,name=name,proto3" json:"name,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Icon          map[string]string `protobuf:"bytes,9,rep,name=icon,proto3" json:"icon,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type GoodsEffect struct {
	Type          int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Online        string `protobuf:"bytes,2,opt,name=online,proto3" json:"online,omitempty"`
	ApiKey        string `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret     string `protobuf:"bytes,4,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	EffectAddress string `protobuf:"bytes,5,opt,name=effect_address,json=effectAddress,proto3" json:"effect_address,omitempty"`
	StyleId       string `protobuf:"bytes,6,opt,name=style_id,json=styleId,proto3" json:"style_id,omitempty"`
}

type GoodsFileRDB struct {
	FileSize string `protobuf:"bytes,1,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	FileUrl  string `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	FileId   string `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	MId      string `protobuf:"bytes,4,opt,name=m_id,json=mId,proto3" json:"m_id,omitempty"`
}
