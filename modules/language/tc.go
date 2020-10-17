package language

import "strings"

var tc = LangSet{
	"managers":  "管理員管理",
	"name":      "用戶名",
	"nickname":  "暱稱",
	"role":      "角色",
	"createdat": "創建時間",
	"updatedat": "更新時間",
	"path":      "路徑",
	"submit":    "提交",
	"filter":    "篩選",

	"new":             "新建",
	"action":          "操作",
	"toggle dropdown": "下拉",
	"delete":          "刪除",
	"refresh":         "刷新",
	"back":            "返回",
	"reset":           "重置",
	"save":            "保存",
	"edit":            "編輯",
	"expand":          "展開",
	"collapse":        "折疊",
	"online":          "在線",
	"setting":         "設置",
	"sign out":        "登出",
	"remove":          "移除",

	"permission manage": "權限管理",
	"menus manage":      "菜單管理",
	"roles manage":      "角色管理",
	"operation log":     "操作日誌",
	"method":            "方法",
	"input":             "輸入",
	"operation":         "操作",
	"menu name":         "菜單名",

	"are you sure to delete": "你確定要刪除嗎",
	"delete succeed":         "刪除成功",
	"yes":                    "確定",
	"got it":                 "知道了",
	"cancel":                 "取消",
	"refresh succeeded":      "刷新成功",

	"avatar":     "頭像",
	"password":   "密碼",
	"slug":       "標誌",
	"permission": "權限",
	"userid":     "用戶ID",
	"content":    "內容",
	"parent":     "父級",
	"icon":       "圖標",
	"uri":        "路徑",

	"export":    "導出",
	"home":      "首頁",
	"all":       "全部",
	"more":      "更多",
	"browse":    "打開",
	"admin":     "管理",
	"users":     "用戶",
	"roles":     "角色",
	"menu":      "菜單",
	"dashboard": "儀表盤",

	"permission denied": "沒有權限",
	"error":             "錯誤",
	"success":           "成功",
	"fail":              "失敗",
	"current page":      "當前頁",

	"confirm":             "確認",
	"edit fail":           "編輯失敗",
	"create fail":         "新增失敗",
	"delete fail":         "刪除失敗",
	"confirm password":    "確認密碼",
	"all method if empty": "為空默認為所有方法",

	"detail":     "詳情",
	"username":   "用戶名",
	"close":      "關閉",
	"login":      "登錄",
	"login fail": "登錄失敗",
	"user":       "用戶",

	"continue editing":  "繼續編輯",
	"continue creating": "繼續新增",

	"goadmin is now running. \nrunning in \"debug\" mode. switch to \"release\" mode in production.\n\n": "GoAdmin 啟動成功。\n目前處於 \"debug\" 模式。請在生產環境中切換為 \"release\" 模式。\n\n",

	"wrong goadmin version, theme %s required goadmin version are %s":    "錯誤的 GoAdmin 版本，當前主題 %s 需要 GoAdmin 版本為 %s",
	"wrong theme version, goadmin %s required version of theme %s is %s": "錯誤的主題版本, GoAdmin %s 需要主題 %s 的版本為 %s",

	"adapter is nil, import the default adapter or use addadapter method add the adapter": "適配器為空，請先 import 對應的適配器或使用 AddAdapter 方法引入",

	"username and password can not be empty":        "用戶名密碼不能為空",
	"operation not allow":                           "不允許的操作",
	"password does not match":                       "密碼不壹致",
	"should be unique":                              "需要保證唯壹",
	"slug exists":                                   "標誌已經存在了",
	"no corresponding options?":                     "沒找到對應選項？",
	"create here.":                                  "在這裏新建壹個。",
	"use for login":                                 "用於登錄",
	"use to display":                                "用來展示",
	"a path a line, without global prefix":          "壹行壹個路徑，換行輸入新路徑，路徑不包含全局路由前綴",
	"slug or http_path or name should not be empty": "標誌或路徑或權限名不能為空",
	"no roles":                                      "無角色",
	"no permission":                                 "沒有權限",
	"fixed the sidebar":                             "固定側邊欄",
	"enter fullscreen":                              "進入全屏",
	"exit fullscreen":                               "退出全屏",
	"wrong captcha":                                 "錯誤的驗證碼",
	"modify success":                                "修改成功",

	"not found":      "找不到記錄",
	"internal error": "系統內部錯誤",
	"unauthorized":   "未認證",

	"login overdue, please login again": "登錄信息過期，請重新登錄",
	"login info":                        "登錄信息",

	"initialize configuration":        "初始化配置",
	"initialize navigation buttons":   "初始化導航欄按鈕",
	"initialize plugins":              "初始化插件",
	"initialize database connections": "初始化數據庫連接",
	"initialize success":              "初始化成功🍺🍺",

	"plugins":          "插件",
	"plugin store":     "插件商店",
	"get more plugins": "獲取更多插件",
	"uninstalled":      "未安裝",
	"plugin setting":   "插件設置",

	"second":  "秒",
	"seconds": "秒",
	"minute":  "分",
	"minutes": "分",
	"hour":    "小時",
	"hours":   "小時",
	"day":     "天",
	"days":    "天",
	"week":    "周",
	"weeks":   "周",
	"month":   "月",
	"months":  "月",
	"year":    "年",
	"years":   "年",

	"config.domain":          "網站域名",
	"config.language":        "網站語言",
	"config.url prefix":      "URL前綴",
	"config.theme":           "主題",
	"config.title":           "標題",
	"config.index url":       "首頁URL",
	"config.login url":       "登錄URL",
	"config.env":             "開發環境",
	"config.color scheme":    "顏色主題",
	"config.cdn url":         "cdn資源URL",
	"config.login title":     "登錄標題",
	"config.auth user table": "登錄用戶表",
	"config.extra":           "額外配置",
	"config.store":           "文件存儲設置",
	"config.databases":       "數據庫設置",
	"config.general":         "通用",
	"config.log":             "日誌",
	"config.site setting":    "網站設置",
	"config.custom":          "定制",
	"config.debug":           "Debug模式",
	"config.site off":        "關閉網站",
	"config.true":            "是",
	"config.false":           "否",

	"config.test":  "測試環境",
	"config.prod":  "生產環境",
	"config.local": "本地環境",

	"config.logo":                        "Logo",
	"config.mini logo":                   "Mini Logo",
	"config.bootstrap file path":         "引导文件路徑",
	"config.go mod file path":            "go.mod文件路徑",
	"config.session life time":           "Session時長",
	"config.custom head html":            "自定義Head HTML",
	"config.custom foot html":            "自定義Foot HTML",
	"config.custom 404 html":             "自定義404頁面",
	"config.custom 403 html":             "自定義403頁面",
	"config.custom 500 html":             "自定義500頁面",
	"config.hide config center entrance": "隱藏配置中心入口",
	"config.hide app info entrance":      "隱藏應用信息入口",
	"config.hide tool entrance":          "隱藏工具入口",
	"config.hide plugin entrance":        "隱藏插件列表入口",
	"config.footer info":                 "自定義底部信息",
	"config.login logo":                  "登錄Logo",
	"config.no limit login ip":           "取消限制多IP登錄",
	"config.operation log off":           "關閉操作日誌",
	"config.allow delete operation log":  "允許刪除操作日誌",
	"config.animation type":              "動畫類型",
	"config.animation duration":          "動畫間隔（秒）",
	"config.animation delay":             "動畫延遲（秒）",
	"config.file upload engine":          "文件上傳引擎",

	"config.logger rotate":             "日誌切割設置",
	"config.logger rotate max size":    "存儲最大文件大小（m）",
	"config.logger rotate max backups": "存儲最多文件數",
	"config.logger rotate max age":     "最長存儲時間（天）",
	"config.logger rotate compress":    "壓縮",

	"config.info log path":         "信息日誌存儲路徑",
	"config.error log path":        "錯誤日誌存儲路徑",
	"config.access log path":       "訪問日誌存儲路徑",
	"config.info log off":          "關閉信息日誌",
	"config.error log off":         "關閉錯誤日誌",
	"config.access log off":        "關閉訪問日誌",
	"config.access assets log off": "關閉靜態資源訪問日誌",
	"config.sql log on":            "打開SQL日誌",
	"config.log level":             "日誌級別",

	"config.logger rotate encoder":                "日誌encoder設置",
	"config.logger rotate encoder time key":       "Time Key",
	"config.logger rotate encoder level key":      "Level Key",
	"config.logger rotate encoder name key":       "Name Key",
	"config.logger rotate encoder caller key":     "Caller Key",
	"config.logger rotate encoder message key":    "Message Key",
	"config.logger rotate encoder stacktrace key": "Stacktrace Key",
	"config.logger rotate encoder level":          "Level字段編碼",
	"config.logger rotate encoder time":           "Time字段編碼",
	"config.logger rotate encoder duration":       "Duration字段編碼",
	"config.logger rotate encoder caller":         "Caller字段編碼",
	"config.logger rotate encoder encoding":       "輸出格式",

	"config.capital":        "大寫",
	"config.capitalcolor":   "大寫帶顏色",
	"config.lowercase":      "小寫",
	"config.lowercasecolor": "小寫帶顏色",

	"config.seconds":     "秒",
	"config.nanosecond":  "納秒",
	"config.microsecond": "微秒",
	"config.millisecond": "毫秒",

	"config.full path":  "完整路徑",
	"config.short path": "簡短路徑",

	"config.do not modify when you have not set up all assets": "不要修改，當妳還沒有設置好所有資源文件的時候",
	"config.it will work when theme is adminlte":               "當主題為adminlte時生效",
	"config.must bigger than 900 seconds":                      "必須大於900秒",

	"config.language." + CN:                  "中文",
	"config.language." + EN:                  "英文",
	"config.language." + JP:                  "日文",
	"config.language." + strings.ToLower(TC): "繁體中文",

	"config.modify site config":         "修改網站配置",
	"config.modify site config success": "修改網站配置成功",
	"config.modify site config fail":    "修改網站配置失敗",

	"system.system info":     "應用系統信息",
	"system.application":     "應用信息",
	"system.application run": "應用運行信息",
	"system.system":          "系統信息",

	"system.process_id":                           "進程ID",
	"system.golang_version":                       "Golang版本",
	"system.server_uptime":                        "服務運行時間",
	"system.current_goroutine":                    "當前 Goroutines 數量",
	"system.current_memory_usage":                 "當前內存使用量",
	"system.total_memory_allocated":               "所有被分配的內存",
	"system.memory_obtained":                      "內存占用量",
	"system.pointer_lookup_times":                 "指針查找次數",
	"system.memory_allocate_times":                "內存分配次數",
	"system.memory_free_times":                    "內存釋放次數",
	"system.current_heap_usage":                   "當前 Heap 內存使用量",
	"system.heap_memory_obtained":                 "Heap 內存占用量",
	"system.heap_memory_idle":                     "Heap 內存空閑量",
	"system.heap_memory_in_use":                   "正在使用的 Heap 內存",
	"system.heap_memory_released":                 "被釋放的 Heap 內存",
	"system.heap_objects":                         "Heap 對象數量",
	"system.bootstrap_stack_usage":                "啟動 Stack 使用量",
	"system.stack_memory_obtained":                "被分配的 Stack 內存",
	"system.mspan_structures_usage":               "MSpan 結構內存使用量",
	"system.mspan_structures_obtained":            "被分配的 MSpan 結構內存",
	"system.mcache_structures_usage":              "MCache 結構內存使用量",
	"system.mcache_structures_obtained":           "被分配的 MCache 結構內存",
	"system.profiling_bucket_hash_table_obtained": "被分配的剖析哈希表內存",
	"system.gc_metadata_obtained":                 "被分配的 GC 元數據內存",
	"system.other_system_allocation_obtained":     "其它被分配的系統內存",
	"system.next_gc_recycle":                      "下次 GC 內存回收量",
	"system.last_gc_time":                         "距離上次 GC 時間",
	"system.total_gc_time":                        "GC 執行時間總量",
	"system.total_gc_pause":                       "GC 暫停時間總量",
	"system.last_gc_pause":                        "上次 GC 暫停時間",
	"system.gc_times":                             "GC 執行次數",

	"system.cpu_logical_core": "cpu邏輯核數",
	"system.cpu_core":         "cpu物理核數",
	"system.os_platform":      "系統平臺",
	"system.os_family":        "系統家族",
	"system.os_version":       "系統版本",
	"system.load1":            "1分鐘內負載",
	"system.load5":            "5分鐘內負載",
	"system.load15":           "15分鐘內負載",
	"system.mem_total":        "總內存",
	"system.mem_available":    "可用內存",
	"system.mem_used":         "使用內存",

	"system.app_name":         "應用名",
	"system.go_admin_version": "應用版本",
	"system.theme_name":       "主題",
	"system.theme_version":    "主題版本",

	"tool.tool":                   "工具",
	"tool.table":                  "表格",
	"tool.connection":             "連接",
	"tool.package":                "包名",
	"tool.output":                 "輸出路徑",
	"tool.output path is empty":   "輸出路徑為空",
	"tool.field":                  "字段名",
	"tool.title":                  "標題",
	"tool.field name":             "字段名",
	"tool.db type":                "數據類型",
	"tool.form type":              "表單類型",
	"tool.generate table model":   "生成CRUD模型",
	"tool.primarykey":             "主鍵",
	"tool.field filterable":       "可篩選",
	"tool.field sortable":         "可排序",
	"tool.yes":                    "是",
	"tool.no":                     "否",
	"tool.generate success":       "生成成功",
	"tool.hide filter area":       "隱藏篩選框",
	"tool.use absolute path":      "使用絕對路徑",
	"tool.hide":                   "隱藏",
	"tool.show":                   "顯示",
	"tool.display":                "顯示",
	"tool.basic info":             "基本信息",
	"tool.table info":             "表格信息",
	"tool.form info":              "表單信息",
	"tool.field editable":         "允許編輯",
	"tool.field can add":          "允許新增",
	"tool.info field editable":    "可編輯",
	"tool.extra import package":   "導入包",
	"tool.field default":          "默認值",
	"tool.filter area":            "篩選框",
	"tool.new button":             "新建按鈕",
	"tool.export button":          "導出按鈕",
	"tool.edit button":            "編輯按鈕",
	"tool.delete button":          "刪除按鈕",
	"tool.detail button":          "詳情按鈕",
	"tool.filter button":          "篩選按鈕",
	"tool.row selector":           "列選擇按鈕",
	"tool.pagination":             "分頁",
	"tool.query info":             "查詢信息",
	"tool.filter form layout":     "篩選表單布局",
	"tool.continue edit checkbox": "繼續編輯按鈕",
	"tool.continue new checkbox":  "繼續新增按鈕",
	"tool.reset button":           "重設按鈕",
	"tool.back button":            "返回按鈕",
	"tool.generate":               "生成",
	"tool.generated tables":       "生成過的表格",
	"tool.description":            "描述",
	"tool.label":                  "標簽",
	"tool.image":                  "圖片",
	"tool.bool":                   "布爾",
	"tool.link":                   "鏈接",
	"tool.fileSize":               "文件大小",
	"tool.date":                   "日期",
	"tool.icon":                   "Icon",
	"tool.dot":                    "標點",
	"tool.progressBar":            "進度條",
	"tool.loading":                "Loading",
	"tool.downLoadable":           "可下載",
	"tool.copyable":               "可復制",
	"tool.carousel":               "圖片輪播",
	"tool.qrcode":                 "二維碼",
	"tool.field hide":             "隱藏",
	"tool.field display":          "顯示",
	"tool.table permission":       "生成表格權限",
	"tool.extra code":             "額外代碼",

	"tool.field display normal":     "顯示",
	"tool.field diplay hide":        "隱藏",
	"tool.field diplay edit hide":   "編輯隱藏",
	"tool.field diplay create hide": "新建隱藏",

	"tool.detail display":             "顯示",
	"tool.detail info":                "詳情頁信息",
	"tool.follow list page":           "跟隨列表頁",
	"tool.inherit from list page":     "繼承列表頁",
	"tool.independent from list page": "獨立",

	"tool.generate table model success": "生成成功",
	"tool.generate table model fail":    "生成失敗",

	"generator.query":                 "查詢",
	"generator.show edit form page":   "編輯頁顯示",
	"generator.show create form page": "新建記錄頁顯示",
	"generator.edit":                  "編輯",
	"generator.create":                "新建",
	"generator.delete":                "刪除",
	"generator.export":                "導出",

	"plugin.plugin":                         "插件",
	"plugin.plugin detail":                  "插件詳情",
	"plugin.introduction":                   "介紹",
	"plugin.website":                        "網站",
	"plugin.version":                        "版本",
	"plugin.created at":                     "創建日期",
	"plugin.updated at":                     "更新日期",
	"plugin.provided by %s":                 "由 %s 提供",
	"plugin.upgrade":                        "升級",
	"plugin.install":                        "安裝",
	"plugin.info":                           "詳細信息",
	"plugin.download":                       "下载",
	"plugin.buy":                            "購買",
	"plugin.downloading":                    "下载中",
	"plugin.login":                          "登錄",
	"plugin.login to goadmin member system": "登錄到GoAdmin會員系統",
	"plugin.account":                        "賬戶名",
	"plugin.password":                       "密碼",
	"plugin.learn more":                     "了解更多",

	"plugin.no account? click %s here %s to register.":    "沒有賬號？點擊%s這裏%s註冊。",
	"plugin.download fail, wrong name":                    "下載失敗，錯誤的參數",
	"plugin.change to debug mode first":                   "先切換到debug模式",
	"plugin.download fail, plugin not exist":              "下載失敗，插件不存在",
	"plugin.download fail":                                "下載失敗",
	"plugin.golang develop environment does not exist":    "golang開發環境不存在",
	"plugin.download success, restart to install":         "下載成功，重啟程序進行安裝",
	"plugin.restart to install":                           "重啟程序進行安裝",
	"plugin.can not connect to the goadmin remote server": "不能連接到GoAdmin遠程服務器，請檢查您的網絡連接。",

	"admin.basic admin": "基礎Admin",
	"admin.a built-in plugins of goadmin which help you to build a crud manager platform quickly.": "壹個內置GoAdmin插件，幫助您快速搭建curd簡易管理後臺。",
	"admin.official": "GoAdmin官方",
}
