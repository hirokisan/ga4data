package ga4data

const (
	// DimensionAchievementID : The achievement ID in a game for an event. Populated by the event parameter 'achievement_id'.
	DimensionAchievementID = "achievementId"
	// DimensionAdFormat : Describes the way ads looked and where they were located. Typical formats include 'Interstitial', 'Banner', 'Rewarded', and 'Native advanced'.
	DimensionAdFormat = "adFormat"
	// DimensionAdSourceName : The source network that served the ad. Typical sources include 'AdMob Network', 'Liftoff', 'Facebook Audience Network', and 'Mediated house ads'.
	DimensionAdSourceName = "adSourceName"
	// DimensionAdUnitName : The name you chose to describe this Ad unit. Ad units are containers you place in your apps to show ads to users.
	DimensionAdUnitName = "adUnitName"
	// DimensionAppVersion : The app's versionName (Android) or short bundle version (iOS).
	DimensionAppVersion = "appVersion"
	// DimensionAudienceID : The numeric identifier of an Audience. Users are reported in the audiences to which they belonged during the report's date range. Current user behavior does not affect historical audience membership in reports.
	DimensionAudienceID = "audienceId"
	// DimensionAudienceName : The given name of an Audience. Users are reported in the audiences to which they belonged during the report's date range. Current user behavior does not affect historical audience membership in reports.
	DimensionAudienceName = "audienceName"
	// DimensionBrandingInterest : Interests demonstrated by users who are higher in the shopping funnel. Users can be counted in multiple interest categories. For example, 'Shoppers', 'Lifestyles & Hobbies/Pet Lovers', or 'Travel/Travel Buffs/Beachbound Travelers'.
	DimensionBrandingInterest = "brandingInterest"
	// DimensionBrowser : The browsers used to view your website.
	DimensionBrowser = "browser"
	// DimensionCampaignID : The identifier of the marketing campaign. Present only for conversion events. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionCampaignID = "campaignId"
	// DimensionCampaignName : The name of the marketing campaign. Present only for conversion events. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionCampaignName = "campaignName"
	// DimensionCharacter : The player character in a game for an event. Populated by the event parameter 'character'.
	DimensionCharacter = "character"
	// DimensionCity : The city from which the user activity originated.
	DimensionCity = "city"
	// DimensionCityID : The geographic ID of the city from which the user activity originated, derived from their IP address.
	DimensionCityID = "cityId"
	// DimensionCohort : The cohort's name in the request. A cohort is a set of users who started using your website or app in any consecutive group of days. If a cohort name is not specified in the request, cohorts are named by their zero based index: cohort_0, cohort_1, etc.
	DimensionCohort = "cohort"
	// DimensionCohortNthDay : Day offset relative to the firstSessionDate for the users in the cohort. For example, if a cohort is selected with the start and end date of 2020-03-01, then for the date 2020-03-02, cohortNthDay will be 0001.
	DimensionCohortNthDay = "cohortNthDay"
	// DimensionCohortNthMonth : Month offset relative to the firstSessionDate for the users in the cohort. Month boundaries align with calendar month boundaries. For example, if a cohort is selected with the start and end date in March 2020, then for any date in April 2020, cohortNthMonth will be 0001.
	DimensionCohortNthMonth = "cohortNthMonth"
	// DimensionCohortNthWeek : Week offset relative to the firstSessionDate for the users in the cohort. Weeks start on Sunday and end on Saturday. For example, if a cohort is selected with the start and end date in the range 2020-11-08 to 2020-11-14, then for the dates in the range 2020-11-15 to 2020-11-21, cohortNthWeek will be 0001.
	DimensionCohortNthWeek = "cohortNthWeek"
	// DimensionContentGroup : A category that applies to items of published content. Populated by the event parameter 'content_group'.
	DimensionContentGroup = "contentGroup"
	// DimensionContentID : The identifier of the selected content. Populated by the event parameter 'content_id'.
	DimensionContentID = "contentId"
	// DimensionContentType : The category of the selected content. Populated by the event parameter 'content_type'.
	DimensionContentType = "contentType"
	// DimensionCountry : The country from which the user activity originated.
	DimensionCountry = "country"
	// DimensionCountryID : The geographic ID of the country from which the user activity originated, derived from their IP address. Formatted according to ISO 3166-1 alpha-2 standard.
	DimensionCountryID = "countryId"
	// DimensionDate : The date of the event, formatted as YYYYMMDD.
	DimensionDate = "date"
	// DimensionDateHour : The combined values of date and hour formatted as YYYYMMDDHH.
	DimensionDateHour = "dateHour"
	// DimensionDateHourMinute : The combined values of date, hour, and minute formatted as YYYYMMDDHHMM.
	DimensionDateHourMinute = "dateHourMinute"
	// DimensionDay : The day of the month, a two-digit number from 01 to 31.
	DimensionDay = "day"
	// DimensionDayOfWeek : The day of the week. It returns values in the range [0,6] with Sunday as the first day of the week.
	DimensionDayOfWeek = "dayOfWeek"
	// DimensionDefaultChannelGroup : The conversion's default channel group is based primarily on source and medium. An enumeration which includes 'Direct', 'Organic Search', 'Paid Social', 'Organic Social', 'Email', 'Affiliates', 'Referral', 'Paid Search', 'Video', and 'Display'.
	DimensionDefaultChannelGroup = "defaultChannelGroup"
	// DimensionDeviceCategory : The type of device: Desktop, Tablet, or Mobile.
	DimensionDeviceCategory = "deviceCategory"
	// DimensionDeviceModel : The mobile device model (example: iPhone 10,6).
	DimensionDeviceModel = "deviceModel"
	// DimensionEventName : The name of the event.
	DimensionEventName = "eventName"
	// DimensionFileExtension : The extension of the downloaded file (for example, 'pdf' or 'txt'). Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'file_extension'.
	DimensionFileExtension = "fileExtension"
	// DimensionFileName : The page path of the downloaded file (for example, '/menus/dinner-menu.pdf'). Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'file_name'.
	DimensionFileName = "fileName"
	// DimensionFirstSessionDate : The date the user's first session occurred, formatted as YYYYMMDD.
	DimensionFirstSessionDate = "firstSessionDate"
	// DimensionFirstUserCampaignID : Identifier of the marketing campaign that first acquired the user. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionFirstUserCampaignID = "firstUserCampaignId"
	// DimensionFirstUserCampaignName : Name of the marketing campaign that first acquired the user. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionFirstUserCampaignName = "firstUserCampaignName"
	// DimensionFirstUserDefaultChannelGroup : The default channel group that first acquired the user. Default channel group is based primarily on source and medium. An enumeration which includes 'Direct', 'Organic Search', 'Paid Social', 'Organic Social', 'Email', 'Affiliates', 'Referral', 'Paid Search', 'Video', and 'Display'.
	DimensionFirstUserDefaultChannelGroup = "firstUserDefaultChannelGroup"
	// DimensionFirstUserGoogleAdsAccountName : The Account name from Google Ads that first acquired the user.
	DimensionFirstUserGoogleAdsAccountName = "firstUserGoogleAdsAccountName"
	// DimensionFirstUserGoogleAdsAdGroupID : The Ad Group Id in Google Ads that first acquired the user.
	DimensionFirstUserGoogleAdsAdGroupID = "firstUserGoogleAdsAdGroupId"
	// DimensionFirstUserGoogleAdsAdGroupName : The Ad Group Name in Google Ads that first acquired the user.
	DimensionFirstUserGoogleAdsAdGroupName = "firstUserGoogleAdsAdGroupName"
	// DimensionFirstUserGoogleAdsAdNetworkType : The advertising network that first acquired the user. An enumeration which includes 'Google search', 'Search partners', 'Google Display Network', 'Youtube Search', 'Youtube Videos', 'Cross-network', 'Social', and '(universal campaign)'.
	DimensionFirstUserGoogleAdsAdNetworkType = "firstUserGoogleAdsAdNetworkType"
	// DimensionFirstUserGoogleAdsCampaignID : Identifier of the Google Ads marketing campaign that first acquired the user.
	DimensionFirstUserGoogleAdsCampaignID = "firstUserGoogleAdsCampaignId"
	// DimensionFirstUserGoogleAdsCampaignName : Name of the Google Ads marketing campaign that first acquired the user.
	DimensionFirstUserGoogleAdsCampaignName = "firstUserGoogleAdsCampaignName"
	// DimensionFirstUserGoogleAdsCampaignType : The campaign type of the Google Ads campaign that first acquired the user. Campaign types determine where customers see your ads and the settings and options available to you in Google Ads. Campaign type is an enumeration that includes: Search, Display, Shopping, Video, Discovery, App, Smart, Hotel, Local, and Performance Max. To learn more, see <https://support.google.com/google-ads/answer/2567043>.
	DimensionFirstUserGoogleAdsCampaignType = "firstUserGoogleAdsCampaignType"
	// DimensionFirstUserGoogleAdsCreativeID : The ID of the Google Ads creative that first acquired the user. Creative IDs identify individual ads.
	DimensionFirstUserGoogleAdsCreativeID = "firstUserGoogleAdsCreativeId"
	// DimensionFirstUserGoogleAdsCustomerID : The Customer ID from Google Ads that first acquired the user. Customer IDs in Google Ads uniquely identify Google Ads accounts.
	DimensionFirstUserGoogleAdsCustomerID = "firstUserGoogleAdsCustomerId"
	// DimensionFirstUserGoogleAdsKeyword : The matched keyword that first acquired the user. Keywords are words or phrases describing your product or service that you choose to get your ad in front of the right customers. To learn more about Keywords, see <https://support.google.com/google-ads/answer/6323>.
	DimensionFirstUserGoogleAdsKeyword = "firstUserGoogleAdsKeyword"
	// DimensionFirstUserGoogleAdsQuery : The search query that first acquired the user.
	DimensionFirstUserGoogleAdsQuery = "firstUserGoogleAdsQuery"
	// DimensionFirstUserManualAdContent : The ad content that first acquired the user. Populated by the utm_content parameter.
	DimensionFirstUserManualAdContent = "firstUserManualAdContent"
	// DimensionFirstUserManualTerm : The term that first acquired the user. Populated by the utm_term parameter.
	DimensionFirstUserManualTerm = "firstUserManualTerm"
	// DimensionFirstUserMedium : The medium that first acquired the user to your website or app.
	DimensionFirstUserMedium = "firstUserMedium"
	// DimensionFirstUserSource : The source that first acquired the user to your website or app.
	DimensionFirstUserSource = "firstUserSource"
	// DimensionFirstUserSourceMedium : The combined values of the dimensions 'firstUserSource' and 'firstUserMedium'.
	DimensionFirstUserSourceMedium = "firstUserSourceMedium"
	// DimensionFirstUserSourcePlatform : The source platform that first acquired the user. Please do not depend on this field returning 'Manual' for traffic that uses UTMs; this field will update from returning 'Manual' to returning '(not set)' for an upcoming feature launch.
	DimensionFirstUserSourcePlatform = "firstUserSourcePlatform"
	// DimensionFullPageUrl : The hostname, page path, and query string for web pages visited; for example, the fullPageUrl portion of https://www.example.com/store/contact-us?query_string=true is www.example.com/store/contact-us?query_string=true.
	DimensionFullPageUrl = "fullPageUrl"
	// DimensionGoogleAdsAccountName : The Account name from Google Ads for the campaign that led to the conversion event. Corresponds to customer.descriptive_name in the Google Ads API.
	DimensionGoogleAdsAccountName = "googleAdsAccountName"
	// DimensionGoogleAdsAdGroupID : The ad group id attributed to the conversion event.
	DimensionGoogleAdsAdGroupID = "googleAdsAdGroupId"
	// DimensionGoogleAdsAdGroupName : The ad group name attributed to the conversion event.
	DimensionGoogleAdsAdGroupName = "googleAdsAdGroupName"
	// DimensionGoogleAdsAdNetworkType : The advertising network type of the conversion. An enumeration which includes 'Google search', 'Search partners', 'Google Display Network', 'Youtube Search', 'Youtube Videos', 'Cross-network', 'Social', and '(universal campaign)'.
	DimensionGoogleAdsAdNetworkType = "googleAdsAdNetworkType"
	// DimensionGoogleAdsCampaignID : The campaign ID for the Google Ads campaign attributed to the conversion event.
	DimensionGoogleAdsCampaignID = "googleAdsCampaignId"
	// DimensionGoogleAdsCampaignName : The campaign name for the Google Ads campaign attributed to the conversion event.
	DimensionGoogleAdsCampaignName = "googleAdsCampaignName"
	// DimensionGoogleAdsCampaignType : The campaign type for the Google Ads campaign attributed to the conversion event. Campaign types determine where customers see your ads and the settings and options available to you in Google Ads. Campaign type is an enumeration that includes: Search, Display, Shopping, Video, Discovery, App, Smart, Hotel, Local, and Performance Max. To learn more, see <https://support.google.com/google-ads/answer/2567043>.
	DimensionGoogleAdsCampaignType = "googleAdsCampaignType"
	// DimensionGoogleAdsCreativeID : The ID of the Google Ads creative attributed to the conversion event. Creative IDs identify individual ads.
	DimensionGoogleAdsCreativeID = "googleAdsCreativeId"
	// DimensionGoogleAdsCustomerID : The Customer ID from Google Ads for the campaign that led to conversion event. Customer IDs in Google Ads uniquely identify Google Ads accounts.
	DimensionGoogleAdsCustomerID = "googleAdsCustomerId"
	// DimensionGoogleAdsKeyword : The matched keyword that led to the conversion event. Keywords are words or phrases describing your product or service that you choose to get your ad in front of the right customers. To learn more about Keywords, see <https://support.google.com/google-ads/answer/6323>.
	DimensionGoogleAdsKeyword = "googleAdsKeyword"
	// DimensionGoogleAdsQuery : The search query that led to the conversion event.
	DimensionGoogleAdsQuery = "googleAdsQuery"
	// DimensionGroupID : The player group ID in a game for an event. Populated by the event parameter 'group_id'.
	DimensionGroupID = "groupId"
	// DimensionHostName : Includes the subdomain and domain names of a URL; for example, the Host Name of www.example.com/contact.html is www.example.com.
	DimensionHostName = "hostName"
	// DimensionHour : The two-digit hour of the day that the event was logged. This dimension ranges from 0-23 and is reported in your property's timezone.
	DimensionHour = "hour"
	// DimensionIsConversionEvent : The string 'true' if the event is a conversion. Events are marked as conversions at collection time; changes to an event's conversion marking apply going forward. You can mark any event as a conversion in Google Analytics, and some events (i.e. first_open, purchase) are marked as conversions by default. To learn more, see <https://support.google.com/analytics/answer/9267568>.
	DimensionIsConversionEvent = "isConversionEvent"
	// DimensionItemAffiliation : The name or code of the affiliate (partner/vendor; if any) associated with an individual item. Populated by the 'affiliation' item parameter.
	DimensionItemAffiliation = "itemAffiliation"
	// DimensionItemBrand : Brand name of the item.
	DimensionItemBrand = "itemBrand"
	// DimensionItemCategory : The hierarchical category in which the item is classified. For example, in Apparel/Mens/Summer/Shirts/T-shirts, Apparel is the item category.
	DimensionItemCategory = "itemCategory"
	// DimensionItemCategory2 : The hierarchical category in which the item is classified. For example, in Apparel/Mens/Summer/Shirts/T-shirts, Mens is the item category 2.
	DimensionItemCategory2 = "itemCategory2"
	// DimensionItemCategory3 : The hierarchical category in which the item is classified. For example, in Apparel/Mens/Summer/Shirts/T-shirts, Summer is the item category 3.
	DimensionItemCategory3 = "itemCategory3"
	// DimensionItemCategory4 : The hierarchical category in which the item is classified. For example, in Apparel/Mens/Summer/Shirts/T-shirts, Shirts is the item category 4.
	DimensionItemCategory4 = "itemCategory4"
	// DimensionItemCategory5 : The hierarchical category in which the item is classified. For example, in Apparel/Mens/Summer/Shirts/T-shirts, T-shirts is the item category 5.
	DimensionItemCategory5 = "itemCategory5"
	// DimensionItemID : The ID of the item.
	DimensionItemID = "itemId"
	// DimensionItemListID : The ID of the item list.
	DimensionItemListID = "itemListId"
	// DimensionItemListName : The name of the item list.
	DimensionItemListName = "itemListName"
	// DimensionItemName : The name of the item.
	DimensionItemName = "itemName"
	// DimensionItemPromotionCreativeName : The name of the item-promotion creative.
	DimensionItemPromotionCreativeName = "itemPromotionCreativeName"
	// DimensionItemPromotionID : The ID of the item promotion.
	DimensionItemPromotionID = "itemPromotionId"
	// DimensionItemPromotionName : The name of the promotion for the item.
	DimensionItemPromotionName = "itemPromotionName"
	// DimensionItemVariant : The specific variation of a product. e.g., XS, S, M, L for size; or Red, Blue, Green, Black for color. Populated by the 'item_variant' parameter.
	DimensionItemVariant = "itemVariant"
	// DimensionLandingPage : The page path + query string associated with the first pageview in a session.
	DimensionLandingPage = "landingPage"
	// DimensionLanguage : The language setting of the user's browser or device. e.g. English
	DimensionLanguage = "language"
	// DimensionLanguageCode : The language setting (ISO 639) of the user's browser or device. e.g. 'en-us'
	DimensionLanguageCode = "languageCode"
	// DimensionLevel : The player's level in a game. Populated by the event parameter 'level'.
	DimensionLevel = "level"
	// DimensionLinkClasses : The HTML class attribute for an outbound link. For example if a user clicks a link '<a class="center" href="www.youtube.com">', this dimension will return 'center'. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'link_classes'.
	DimensionLinkClasses = "linkClasses"
	// DimensionLinkDomain : The destination domain of the outbound link. For example if a user clicks a link '<a href="www.youtube.com">', this dimension will return 'youtube.com'. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'link_domain'.
	DimensionLinkDomain = "linkDomain"
	// DimensionLinkID : The HTML id attribute for an outbound link or file download. For example if a user clicks a link '<a id="socialLinks" href="www.youtube.com">', this dimension will return 'socialLinks'. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'link_id'.
	DimensionLinkID = "linkId"
	// DimensionLinkText : The link text of the file download. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'link_text'.
	DimensionLinkText = "linkText"
	// DimensionLinkUrl : The full url for an outbound link or file download. For example if a user clicks a link '<a href="https://www.youtube.com/results?search_query=analytics">', this dimension will return 'https://www.youtube.com/results?search_query=analytics'. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'link_url'.
	DimensionLinkUrl = "linkUrl"
	// DimensionManualAdContent : The ad content attributed to the conversion event. Populated by the utm_content parameter.
	DimensionManualAdContent = "manualAdContent"
	// DimensionManualTerm : The term attributed to the conversion event. Populated by the utm_term parameter.
	DimensionManualTerm = "manualTerm"
	// DimensionMedium : The medium attributed to the conversion event.
	DimensionMedium = "medium"
	// DimensionMethod : The method by which an event was triggered. Populated by the event parameter 'method'.
	DimensionMethod = "method"
	// DimensionMinute : The two-digit minute of the hour that the event was logged. This dimension ranges from 0-59 and is reported in your property's timezone.
	DimensionMinute = "minute"
	// DimensionMobileDeviceBranding : Manufacturer or branded name (examples: Samsung, HTC, Verizon, T-Mobile).
	DimensionMobileDeviceBranding = "mobileDeviceBranding"
	// DimensionMobileDeviceMarketingName : The branded device name (examples: Galaxy S10 or P30 Pro).
	DimensionMobileDeviceMarketingName = "mobileDeviceMarketingName"
	// DimensionMobileDeviceModel : The mobile device model name (examples: iPhone X or SM-G950F).
	DimensionMobileDeviceModel = "mobileDeviceModel"
	// DimensionMonth : The month of the event, a two digit integer from 01 to 12.
	DimensionMonth = "month"
	// DimensionNewVsReturning : New users have 0 previous sessions, and returning users have 1 or more previous sessions. This dimension returns two values: 'new' or 'returning'.
	DimensionNewVsReturning = "newVsReturning"
	// DimensionNthDay : The number of days since the start of the date range.
	DimensionNthDay = "nthDay"
	// DimensionNthHour : The number of hours since the start of the date range. The starting hour is 0000.
	DimensionNthHour = "nthHour"
	// DimensionNthMinute : The number of minutes since the start of the date range. The starting minute is 0000.
	DimensionNthMinute = "nthMinute"
	// DimensionNthMonth : The number of months since the start of a date range. The starting month is 0000.
	DimensionNthMonth = "nthMonth"
	// DimensionNthWeek : A number representing the number of weeks since the start of a date range.
	DimensionNthWeek = "nthWeek"
	// DimensionNthYear : The number of years since the start of the date range. The starting year is 0000.
	DimensionNthYear = "nthYear"
	// DimensionOperatingSystem : The operating systems used by visitors to your app or website. Includes desktop and mobile operating systems such as Windows and Android.
	DimensionOperatingSystem = "operatingSystem"
	// DimensionOperatingSystemVersion : The operating system versions used by visitors to your website or app. For example, Android 10's version is 10, and iOS 13.5.1's version is 13.5.1.
	DimensionOperatingSystemVersion = "operatingSystemVersion"
	// DimensionOperatingSystemWithVersion : The operating system and version. For example, Android 10 or Windows 7.
	DimensionOperatingSystemWithVersion = "operatingSystemWithVersion"
	// DimensionOrderCoupon : Code for the order-level coupon.
	DimensionOrderCoupon = "orderCoupon"
	// DimensionOutbound : Returns 'true' if the link lead to a site is not a part of the property’s domain. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'outbound'.
	DimensionOutbound = "outbound"
	// DimensionPageLocation : The protocol, hostname, page path, and query string for web pages visited; for example, the pageLocation portion of https://www.example.com/store/contact-us?query_string=true is https://www.example.com/store/contact-us?query_string=true. Populated by the event parameter 'page_location'.
	DimensionPageLocation = "pageLocation"
	// DimensionPagePath : The portion of the URL between the hostname and query string for web pages visited; for example, the pagePath portion of https://www.example.com/store/contact-us?query_string=true is /store/contact-us.
	DimensionPagePath = "pagePath"
	// DimensionPagePathPlusQueryString : The portion of the URL following the hostname for web pages visited; for example, the pagePathPlusQueryString portion of https://www.example.com/store/contact-us?query_string=true is /store/contact-us?query_string=true.
	DimensionPagePathPlusQueryString = "pagePathPlusQueryString"
	// DimensionPageReferrer : The full referring URL including the hostname and path. This referring URL is the user's previous URL and can be this website's domain or other domains. Populated by the event parameter 'page_referrer'.
	DimensionPageReferrer = "pageReferrer"
	// DimensionPageTitle : The web page titles used on your site.
	DimensionPageTitle = "pageTitle"
	// DimensionPercentScrolled : The percentage down the page that the user has scrolled (for example, '90'). Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'percent_scrolled'.
	DimensionPercentScrolled = "percentScrolled"
	// DimensionPlatform : The platform on which your app or website ran; for example, web, iOS, or Android. To determine a stream's type in a report, use both platform and streamId.
	DimensionPlatform = "platform"
	// DimensionPlatformDeviceCategory : The platform and type of device on which your website or mobile app ran. (example: Android / mobile)
	DimensionPlatformDeviceCategory = "platformDeviceCategory"
	// DimensionRegion : The geographic region from which the user activity originated, derived from their IP address.
	DimensionRegion = "region"
	// DimensionScreenResolution : The screen resolution of the user's monitor. For example, 1920x1080.
	DimensionScreenResolution = "screenResolution"
	// DimensionSearchTerm : The term searched by the user. For example if the user visits '/some-page.html?q=some-term', this dimension returns 'some-term'. Automatically populated if Enhanced Measurement is enabled. Populated by the event parameter 'search_term'.
	DimensionSearchTerm = "searchTerm"
	// DimensionSessionCampaignID : The marketing campaign id for a session. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionSessionCampaignID = "sessionCampaignId"
	// DimensionSessionCampaignName : The marketing campaign name for a session. Includes Google Ads Campaigns, Manual Campaigns, & other Campaigns.
	DimensionSessionCampaignName = "sessionCampaignName"
	// DimensionSessionDefaultChannelGroup : The session's default channel group is based primarily on source and medium. An enumeration which includes 'Direct', 'Organic Search', 'Paid Social', 'Organic Social', 'Email', 'Affiliates', 'Referral', 'Paid Search', 'Video', and 'Display'.
	DimensionSessionDefaultChannelGroup = "sessionDefaultChannelGroup"
	// DimensionSessionGoogleAdsAccountName : The Account name from Google Ads that led to the session. Corresponds to customer.descriptive_name in the Google Ads API.
	DimensionSessionGoogleAdsAccountName = "sessionGoogleAdsAccountName"
	// DimensionSessionGoogleAdsAdGroupID : The Ad Group Id in Google Ads for a session.
	DimensionSessionGoogleAdsAdGroupID = "sessionGoogleAdsAdGroupId"
	// DimensionSessionGoogleAdsAdGroupName : The Ad Group Name in Google Ads for a session.
	DimensionSessionGoogleAdsAdGroupName = "sessionGoogleAdsAdGroupName"
	// DimensionSessionGoogleAdsAdNetworkType : The advertising network that led to the session. An enumeration which includes 'Google search', 'Search partners', 'Google Display Network', 'Youtube Search', 'Youtube Videos', 'Cross-network', 'Social', and '(universal campaign)'.
	DimensionSessionGoogleAdsAdNetworkType = "sessionGoogleAdsAdNetworkType"
	// DimensionSessionGoogleAdsCampaignID : The Campaign ID for the Google Ads Campaign that led to this session.
	DimensionSessionGoogleAdsCampaignID = "sessionGoogleAdsCampaignId"
	// DimensionSessionGoogleAdsCampaignName : The Campaign name for the Google Ads Campaign that led to this session.
	DimensionSessionGoogleAdsCampaignName = "sessionGoogleAdsCampaignName"
	// DimensionSessionGoogleAdsCampaignType : The campaign type for the Google Ads campaign that led to this session. Campaign types determine where customers see your ads and the settings and options available to you in Google Ads. Campaign type is an enumeration that includes: Search, Display, Shopping, Video, Discovery, App, Smart, Hotel, Local, and Performance Max. To learn more, see <https://support.google.com/google-ads/answer/2567043>.
	DimensionSessionGoogleAdsCampaignType = "sessionGoogleAdsCampaignType"
	// DimensionSessionGoogleAdsCreativeID : The ID of the Google Ads creative that lead to a session on your website or app. Creative IDs identify individual ads.
	DimensionSessionGoogleAdsCreativeID = "sessionGoogleAdsCreativeId"
	// DimensionSessionGoogleAdsCustomerID : The Customer ID from Google Ads that led to the session. Customer IDs in Google Ads uniquely identify Google Ads accounts.
	DimensionSessionGoogleAdsCustomerID = "sessionGoogleAdsCustomerId"
	// DimensionSessionGoogleAdsKeyword : The matched keyword that led to the session. Keywords are words or phrases describing your product or service that you choose to get your ad in front of the right customers. To learn more about Keywords, see <https://support.google.com/google-ads/answer/6323>.
	DimensionSessionGoogleAdsKeyword = "sessionGoogleAdsKeyword"
	// DimensionSessionGoogleAdsQuery : The search query that led to the session.
	DimensionSessionGoogleAdsQuery = "sessionGoogleAdsQuery"
	// DimensionSessionManualAdContent : The ad content that led to a session. Populated by the utm_content parameter.
	DimensionSessionManualAdContent = "sessionManualAdContent"
	// DimensionSessionManualTerm : The term that led to a session. Populated by the utm_term parameter.
	DimensionSessionManualTerm = "sessionManualTerm"
	// DimensionSessionMedium : The medium that initiated a session on your website or app.
	DimensionSessionMedium = "sessionMedium"
	// DimensionSessionSa360AdGroupName : The Ad Group name from Search Ads 360 that led to this session.
	DimensionSessionSa360AdGroupName = "sessionSa360AdGroupName"
	// DimensionSessionSa360CampaignID : The Campaign ID from Search Ads 360 that led to this session.
	DimensionSessionSa360CampaignID = "sessionSa360CampaignId"
	// DimensionSessionSa360CampaignName : The Campaign name from Search Ads 360 that led to this session.
	DimensionSessionSa360CampaignName = "sessionSa360CampaignName"
	// DimensionSessionSa360CreativeFormat : The type of creative in Search Ads 360 that led to this session. For example, 'Responsive search ad' or 'Expanded text ad'. To learn more, see [GA4 Traffic Source Dimensions](https://support.google.com/analytics/answer/9143382#traffic-source).
	DimensionSessionSa360CreativeFormat = "sessionSa360CreativeFormat"
	// DimensionSessionSa360EngineAccountID : The ID of the engine account in SA360 that led to this session.
	DimensionSessionSa360EngineAccountID = "sessionSa360EngineAccountId"
	// DimensionSessionSa360EngineAccountName : The name of the engine account in SA360 that led to this session.
	DimensionSessionSa360EngineAccountName = "sessionSa360EngineAccountName"
	// DimensionSessionSa360EngineAccountType : The type of the engine account in Search Ads 360 that led to this session. For example, 'google ads', 'bing', or 'baidu'.
	DimensionSessionSa360EngineAccountType = "sessionSa360EngineAccountType"
	// DimensionSessionSa360Keyword : The search engine keyword from Search Ads 360 that led to this session.
	DimensionSessionSa360Keyword = "sessionSa360Keyword"
	// DimensionSessionSa360Medium : The search engine keyword from Search Ads 360 that led to this session. For example, 'cpc'.
	DimensionSessionSa360Medium = "sessionSa360Medium"
	// DimensionSessionSa360Query : The search query from Search Ads 360 that led to this session.
	DimensionSessionSa360Query = "sessionSa360Query"
	// DimensionSessionSa360Source : The source of the traffic from Search Ads 360 that led to this session. For example, 'example.com' or 'google'.
	DimensionSessionSa360Source = "sessionSa360Source"
	// DimensionSessionSource : The source that initiated a session on your website or app.
	DimensionSessionSource = "sessionSource"
	// DimensionSessionSourceMedium : The combined values of the dimensions 'sessionSource' and 'sessionMedium'.
	DimensionSessionSourceMedium = "sessionSourceMedium"
	// DimensionSessionSourcePlatform : The source platform of the session's campaign. Please do not depend on this field returning 'Manual' for traffic that uses UTMs; this field will update from returning 'Manual' to returning '(not set)' for an upcoming feature launch.
	DimensionSessionSourcePlatform = "sessionSourcePlatform"
	// DimensionShippingTier : The shipping tier (e.g. Ground, Air, Next-day) selected for delivery of the purchased item. Populated by the 'shipping_tier' event parameter.
	DimensionShippingTier = "shippingTier"
	// DimensionSignedInWithUserID : The string 'yes' if the user signed in with the User-ID feature. To learn more about User-ID, see <https://support.google.com/analytics/answer/9213390>.
	DimensionSignedInWithUserID = "signedInWithUserId"
	// DimensionSource : The source attributed to the conversion event.
	DimensionSource = "source"
	// DimensionSourceMedium : The combined values of the dimensions 'source' and 'medium'.
	DimensionSourceMedium = "sourceMedium"
	// DimensionSourcePlatform : The source platform of the conversion event's campaign. Please do not depend on this field returning 'Manual' for traffic that uses UTMs; this field will update from returning 'Manual' to returning '(not set)' for an upcoming feature launch.
	DimensionSourcePlatform = "sourcePlatform"
	// DimensionStreamID : The numeric data stream identifier for your app or website.
	DimensionStreamID = "streamId"
	// DimensionStreamName : The data stream name for your app or website.
	DimensionStreamName = "streamName"
	// DimensionTestDataFilterName : The name of data filters in testing state. You use data filters to include or exclude event data from your reports based on event-parameter values. To learn more, see <https://support.google.com/analytics/answer/10108813>.
	DimensionTestDataFilterName = "testDataFilterName"
	// DimensionTransactionID : The ID of the ecommerce transaction.
	DimensionTransactionID = "transactionId"
	// DimensionUnifiedPagePathScreen : The page path (web) or screen class (app) on which the event was logged.
	DimensionUnifiedPagePathScreen = "unifiedPagePathScreen"
	// DimensionUnifiedPageScreen : The page path and query string (web) or screen class (app) on which the event was logged.
	DimensionUnifiedPageScreen = "unifiedPageScreen"
	// DimensionUnifiedScreenClass : The page title (web) or screen class (app) on which the event was logged.
	DimensionUnifiedScreenClass = "unifiedScreenClass"
	// DimensionUnifiedScreenName : The page title (web) or screen name (app) on which the event was logged.
	DimensionUnifiedScreenName = "unifiedScreenName"
	// DimensionUserAgeBracket : User age brackets.
	DimensionUserAgeBracket = "userAgeBracket"
	// DimensionUserGender : User gender.
	DimensionUserGender = "userGender"
	// DimensionVideoProvider : The source of the video (for example, 'youtube'). Automatically populated for embedded videos if Enhanced Measurement is enabled. Populated by the event parameter 'video_provider'.
	DimensionVideoProvider = "videoProvider"
	// DimensionVideoTitle : The title of the video. Automatically populated for embedded videos if Enhanced Measurement is enabled. Populated by the event parameter 'video_title'.
	DimensionVideoTitle = "videoTitle"
	// DimensionVideoUrl : The url of the video. Automatically populated for embedded videos if Enhanced Measurement is enabled. Populated by the event parameter 'video_url'.
	DimensionVideoUrl = "videoUrl"
	// DimensionVirtualCurrencyName : The name of a virtual currency with which the user is interacting. i.e. spending or purchasing gems in a game. Populated by the 'virtual_currency_name' event parameter.
	DimensionVirtualCurrencyName = "virtualCurrencyName"
	// DimensionVisible : Returns 'true' if the content is visible. Automatically populated for embedded videos if Enhanced Measurement is enabled. Populated by the event parameter 'visible'.
	DimensionVisible = "visible"
	// DimensionWeek : The week of the event, a two-digit number from 01 to 53. Each week starts on Sunday. January 1st is always in week 01. The first and last week of the year have fewer than 7 days in most years. Weeks other than the first and the last week of the year always have 7 days. For years where January 1st is a Sunday, the first week of that year and the last week of the prior year have 7 days.
	DimensionWeek = "week"
	// DimensionYear : The four-digit year of the event e.g. 2020.
	DimensionYear = "year"
	// MetricActive1DayUsers : The number of distinct active users on your site or app within a 1 day period. The 1 day period includes the last day in the report's date range. Note: this is the same as Active Users.
	MetricActive1DayUsers = "active1DayUsers"
	// MetricActive28DayUsers : The number of distinct active users on your site or app within a 28 day period. The 28 day period includes the last day in the report's date range.
	MetricActive28DayUsers = "active28DayUsers"
	// MetricActive7DayUsers : The number of distinct active users on your site or app within a 7 day period. The 7 day period includes the last day in the report's date range.
	MetricActive7DayUsers = "active7DayUsers"
	// MetricActiveUsers : The number of distinct users who visited your site or app.
	MetricActiveUsers = "activeUsers"
	// MetricAddToCarts : The number of times users added items to their shopping carts.
	MetricAddToCarts = "addToCarts"
	// MetricAdUnitExposure : The time that an ad unit was exposed to a user, in milliseconds.
	MetricAdUnitExposure = "adUnitExposure"
	// MetricAdvertiserAdClicks : Total number of times users have clicked on an ad to reach the property. Includes clicks from linked integrations like linked Search Ads 360 advertisers. Also includes uploaded clicks from data import.
	MetricAdvertiserAdClicks = "advertiserAdClicks"
	// MetricAdvertiserAdCost : The total amount you paid for your ads. Includes costs from linked integrations like linked Google Ads accounts. Also includes uploaded cost from data import; to learn more, see [Import cost data](https://support.google.com/analytics/answer/10071305).
	MetricAdvertiserAdCost = "advertiserAdCost"
	// MetricAdvertiserAdCostPerClick : Ads cost per click is ad cost divided by ad clicks and is often abbreviated CPC.
	MetricAdvertiserAdCostPerClick = "advertiserAdCostPerClick"
	// MetricAdvertiserAdCostPerConversion : Cost per conversion is ad cost divided by conversions.
	MetricAdvertiserAdCostPerConversion = "advertiserAdCostPerConversion"
	// MetricAdvertiserAdImpressions : The total number of impressions. Includes impressions from linked integrations like linked Display & Video 360 advertisers. Also includes uploaded impressions from data import.
	MetricAdvertiserAdImpressions = "advertiserAdImpressions"
	// MetricAveragePurchaseRevenue : The average purchase revenue in the transaction group of events.
	MetricAveragePurchaseRevenue = "averagePurchaseRevenue"
	// MetricAveragePurchaseRevenuePerPayingUser : Average revenue per paying user (ARPPU) is the total purchase revenue per active user that logged a purchase event. The summary metric is for the time period selected.
	MetricAveragePurchaseRevenuePerPayingUser = "averagePurchaseRevenuePerPayingUser"
	// MetricAveragePurchaseRevenuePerUser : The average purchase revenue per active user is the total purchase revenue per active user that logged any event. The summary metric is for the time period selected.
	MetricAveragePurchaseRevenuePerUser = "averagePurchaseRevenuePerUser"
	// MetricAverageRevenuePerUser : Average revenue per active user (ARPU). The summary metric is for the time period selected. ARPU uses Total Revenue and includes AdMob estimated earnings.
	MetricAverageRevenuePerUser = "averageRevenuePerUser"
	// MetricAverageSessionDuration : The average duration (in seconds) of users' sessions.
	MetricAverageSessionDuration = "averageSessionDuration"
	// MetricBounceRate : The percentage of sessions that were not engaged ((Sessions Minus Engaged sessions) divided by Sessions). This metric is returned as a fraction; for example, 0.2761 means 27.61% of sessions were bounces.
	MetricBounceRate = "bounceRate"
	// MetricCartToViewRate : The number of users who added a product(s) to their cart divided by the number of users who viewed the same product(s). This metric is returned as a fraction; for example, 0.1132 means 11.32% of users who viewed a product also added the same product to their cart.
	MetricCartToViewRate = "cartToViewRate"
	// MetricCheckouts : The number of times users started the checkout process. This metric counts the occurrence of the 'begin_checkout' event.
	MetricCheckouts = "checkouts"
	// MetricCohortActiveUsers : The number of users in the cohort who are active in the time window corresponding to the cohort nth day/week/month. For example for the row where cohortNthWeek = 0001, this metric is the number of users (in the cohort) who are active in week 1.
	MetricCohortActiveUsers = "cohortActiveUsers"
	// MetricCohortTotalUsers : The total number of users in the cohort. This metric is the same value in every row of the report for each cohort. Because cohorts are defined by a shared acquisition date, cohortTotalUsers is the same as cohortActiveUsers for the cohort's selection date range. For report rows later than the ochort's selection range, it is typical for cohortActiveUsers to be smaller than cohortTotalUsers. This difference represents users from the cohort that were not active for the later date. cohortTotalUsers is commonly used in the metric expression cohortActiveUsers/cohortTotalUsers to compute a user retention fraction for the cohort. The relationship between activeUsers and totalUsers is not equivalent to the relationship between cohortActiveUsers and cohortTotalUsers.
	MetricCohortTotalUsers = "cohortTotalUsers"
	// MetricConversions : The count of conversion events. Events are marked as conversions at collection time; changes to an event's conversion marking apply going forward. You can mark any event as a conversion in Google Analytics, and some events (i.e. first_open, purchase) are marked as conversions by default. To learn more, see <https://support.google.com/analytics/answer/9267568>.
	MetricConversions = "conversions"
	// MetricCrashAffectedUsers : The number of users that logged a crash in this row of the report. For example if the report is time series by date, this metrics reports total users with at least one crash on this date. Crashes are events with the name "app_exception".
	MetricCrashAffectedUsers = "crashAffectedUsers"
	// MetricCrashFreeUsersRate : The number of users without crash events (in this row of the report) divided by the total number of users. This metric is returned as a fraction; for example, 0.9243 means 92.43% of users were crash-free.
	MetricCrashFreeUsersRate = "crashFreeUsersRate"
	// MetricDauPerMau : The rolling percent of 30-day active users who are also 1-day active users. This metric is returned as a fraction; for example, 0.113 means 11.3% of 30-day active users were also 1-day active users.
	MetricDauPerMau = "dauPerMau"
	// MetricDauPerWau : The rolling percent of 7-day active users who are also 1-day active users. This metric is returned as a fraction; for example, 0.082 means 8.2% of 7-day active users were also 1-day active users.
	MetricDauPerWau = "dauPerWau"
	// MetricEcommercePurchases : The number of times users completed a purchase. This metric counts 'purchase' events; this metric does not count 'in_app_purchase' and subscription events.
	MetricEcommercePurchases = "ecommercePurchases"
	// MetricEngagedSessions : The number of sessions that lasted longer than 10 seconds, or had a conversion event, or had 2 or more screen views.
	MetricEngagedSessions = "engagedSessions"
	// MetricEngagementRate : The percentage of engaged sessions (Engaged sessions divided by Sessions). This metric is returned as a fraction; for example, 0.7239 means 72.39% of sessions were engaged sessions.
	MetricEngagementRate = "engagementRate"
	// MetricEventCount : The count of events.
	MetricEventCount = "eventCount"
	// MetricEventCountPerUser : The average number of events per user (Event count divided by Active users).
	MetricEventCountPerUser = "eventCountPerUser"
	// MetricEventsPerSession : The average number of events per session (Event count divided by Sessions).
	MetricEventsPerSession = "eventsPerSession"
	// MetricEventValue : The sum of the event parameter named 'value'.
	MetricEventValue = "eventValue"
	// MetricFirstTimePurchaserConversionRate : The percentage of active users who made their first purchase. This metric is returned as a fraction; for example, 0.092 means 9.2% of active users were first time purchasers.
	MetricFirstTimePurchaserConversionRate = "firstTimePurchaserConversionRate"
	// MetricFirstTimePurchasers : The number of users that completed their first purchase event.
	MetricFirstTimePurchasers = "firstTimePurchasers"
	// MetricFirstTimePurchasersPerNewUser : The average number of first time purchasers per new user.
	MetricFirstTimePurchasersPerNewUser = "firstTimePurchasersPerNewUser"
	// MetricItemListClickEvents : The number of times users clicked an item when it appeared in a list. This metric counts the occurrence of the 'select_item' event.
	MetricItemListClickEvents = "itemListClickEvents"
	// MetricItemListClickThroughRate : The number of users who selected a list(s) divided by the number of users who viewed the same list(s). This metric is returned as a fraction; for example, 0.2145 means 21.45% of users who viewed a list also selected the same list.
	MetricItemListClickThroughRate = "itemListClickThroughRate"
	// MetricItemListViewEvents : The number of times the item list was viewed. This metric counts the occurrence of the 'view_item_list' event.
	MetricItemListViewEvents = "itemListViewEvents"
	// MetricItemPromotionClickThroughRate : The number of users who selected a promotion(s) divided by the number of users who viewed the same promotion(s). This metric is returned as a fraction; for example, 0.1382 means 13.82% of users who viewed a promotion also selected the promotion.
	MetricItemPromotionClickThroughRate = "itemPromotionClickThroughRate"
	// MetricItemRevenue : The total revenue from items only. Item revenue is the product of its price and quantity. Item revenue excludes tax and shipping values; tax & shipping values are specified at the event and not item level.
	MetricItemRevenue = "itemRevenue"
	// MetricItemsAddedToCart : The number of units added to cart for a single item. This metric counts the quantity of items in 'add_to_cart' events.
	MetricItemsAddedToCart = "itemsAddedToCart"
	// MetricItemsCheckedOut : The number of units checked out for a single item. This metric counts the quantity of items in 'begin_checkout' events.
	MetricItemsCheckedOut = "itemsCheckedOut"
	// MetricItemsClickedInList : The number of units clicked in list for a single item. This metric counts the quantity of items in 'select_item' events.
	MetricItemsClickedInList = "itemsClickedInList"
	// MetricItemsClickedInPromotion : The number of units clicked in promotion for a single item. This metric counts the quantity of items in 'select_promotion' events.
	MetricItemsClickedInPromotion = "itemsClickedInPromotion"
	// MetricItemsPurchased : The number of units for a single item included in purchase events. This metric counts the quantity of items in purchase events.
	MetricItemsPurchased = "itemsPurchased"
	// MetricItemsViewed : The number of units viewed for a single item. This metric counts the quantity of items in 'view_item' events.
	MetricItemsViewed = "itemsViewed"
	// MetricItemsViewedInList : The number of units viewed in list for a single item. This metric counts the quantity of items in 'view_item_list' events.
	MetricItemsViewedInList = "itemsViewedInList"
	// MetricItemsViewedInPromotion : The number of units viewed in promotion for a single item. This metric counts the quantity of items in 'view_promotion' events.
	MetricItemsViewedInPromotion = "itemsViewedInPromotion"
	// MetricItemViewEvents : The number of times the item details were viewed. The metric counts the occurrence of the 'view_item' event.
	MetricItemViewEvents = "itemViewEvents"
	// MetricNewUsers : The number of users who interacted with your site or launched your app for the first time (event triggered: first_open or first_visit).
	MetricNewUsers = "newUsers"
	// MetricOrganicGoogleSearchAveragePosition : The average ranking of your website URLs for the query reported from Search Console. For example, if your site's URL appears at position 3 for one query and position 7 for another query, the average position would be 5 (3+7/2). This metric requires an active Search Console link.
	MetricOrganicGoogleSearchAveragePosition = "organicGoogleSearchAveragePosition"
	// MetricOrganicGoogleSearchClicks : The number of organic Google Search clicks reported from Search Console.  This metric requires an active Search Console link.
	MetricOrganicGoogleSearchClicks = "organicGoogleSearchClicks"
	// MetricOrganicGoogleSearchClickThroughRate : The organic Google Search click through rate reported from Search Console. Click through rate is clicks per impression. This metric is returned as a fraction; for example, 0.0588 means about 5.88% of impressions resulted in a click. This metric requires an active Search Console link.
	MetricOrganicGoogleSearchClickThroughRate = "organicGoogleSearchClickThroughRate"
	// MetricOrganicGoogleSearchImpressions : The number of organic Google Search impressions reported from Search Console. This metric requires an active Search Console link.
	MetricOrganicGoogleSearchImpressions = "organicGoogleSearchImpressions"
	// MetricPromotionClicks : The number of times an item promotion was clicked. This metric counts the occurrence of the 'select_promotion' event.
	MetricPromotionClicks = "promotionClicks"
	// MetricPromotionViews : The number of times an item promotion was viewed. This metric counts the occurrence of the 'view_promotion' event.
	MetricPromotionViews = "promotionViews"
	// MetricPublisherAdClicks : The number of ad_click events.
	MetricPublisherAdClicks = "publisherAdClicks"
	// MetricPublisherAdImpressions : The number of ad_impression events.
	MetricPublisherAdImpressions = "publisherAdImpressions"
	// MetricPurchaserConversionRate : The percentage of active users who made 1 or more purchase transactions. This metric is returned as a fraction; for example, 0.412 means 41.2% of users were purchasers.
	MetricPurchaserConversionRate = "purchaserConversionRate"
	// MetricPurchaseRevenue : The sum of revenue from purchases made in your app or site. Purchase revenue sums the revenue for these events: 'purchase', 'ecommerce_purchase', 'in_app_purchase', 'app_store_subscription_convert', and 'app_store_subscription_renew'. Purchase revenue is specified by the 'value' parameter in tagging.
	MetricPurchaseRevenue = "purchaseRevenue"
	// MetricPurchaseToViewRate : The number of users who purchased a product(s) divided by the number of users who viewed the same product(s). This metric is returned as a fraction; for example, 0.128 means 12.8% of users that viewed a product(s) also purchased the same product(s).
	MetricPurchaseToViewRate = "purchaseToViewRate"
	// MetricReturnOnAdSpend : Return On Ad Spend (ROAS) is total revenue divided by advertiser ad cost.
	MetricReturnOnAdSpend = "returnOnAdSpend"
	// MetricScreenPageViews : The number of app screens or web pages your users viewed. Repeated views of a single page or screen are counted. (screen_view + page_view events).
	MetricScreenPageViews = "screenPageViews"
	// MetricScreenPageViewsPerSession : The number of app screens or web pages your users viewed per session. Repeated views of a single page or screen are counted. (screen_view + page_view events) / sessions.
	MetricScreenPageViewsPerSession = "screenPageViewsPerSession"
	// MetricSessionConversionRate : The percentage of sessions in which any conversion event was triggered.
	MetricSessionConversionRate = "sessionConversionRate"
	// MetricSessions : The number of sessions that began on your site or app (event triggered: session_start).
	MetricSessions = "sessions"
	// MetricSessionsPerUser : The average number of sessions per user (Sessions divided by Active Users).
	MetricSessionsPerUser = "sessionsPerUser"
	// MetricShippingAmount : Shipping amount associated with a transaction. Populated by the 'shipping' event parameter.
	MetricShippingAmount = "shippingAmount"
	// MetricTaxAmount : Tax amount associated with a transaction. Populated by the 'tax' event parameter.
	MetricTaxAmount = "taxAmount"
	// MetricTotalAdRevenue : The total advertising revenue from both Admob and third-party sources.
	MetricTotalAdRevenue = "totalAdRevenue"
	// MetricTotalPurchasers : The number of users that logged purchase events for the time period selected.
	MetricTotalPurchasers = "totalPurchasers"
	// MetricTotalRevenue : The sum of revenue from purchases, subscriptions, and advertising (Purchase revenue plus Subscription revenue plus Ad revenue).
	MetricTotalRevenue = "totalRevenue"
	// MetricTotalUsers : The number of distinct users who have logged at least one event, regardless of whether the site or app was in use when that event was logged.
	MetricTotalUsers = "totalUsers"
	// MetricTransactions : The count of transaction events with purchase revenue. Transaction events are in_app_purchase, ecommerce_purchase, purchase, app_store_subscription_renew, app_store_subscription_convert, and refund.
	MetricTransactions = "transactions"
	// MetricTransactionsPerPurchaser : The average numer of transactions per purchaser.
	MetricTransactionsPerPurchaser = "transactionsPerPurchaser"
	// MetricUserConversionRate : The percentage of users who triggered any conversion event.
	MetricUserConversionRate = "userConversionRate"
	// MetricUserEngagementDuration : The total amount of time (in seconds) your website or app was in the foreground of users' devices.
	MetricUserEngagementDuration = "userEngagementDuration"
	// MetricWauPerMau : The rolling percent of 30-day active users who are also 7-day active users. This metric is returned as a fraction; for example, 0.234 means 23.4% of 30-day active users were also 7-day active users.
	MetricWauPerMau = "wauPerMau"
)
