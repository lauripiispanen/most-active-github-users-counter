package main

type QueryPreset struct {
	include []string
	exclude []string
}

var PRESETS = map[string]QueryPreset{
	"austria": QueryPreset{
		include: []string{"austria", "österreich", "vienna", "wien", "linz", "salzburg", "graz", "innsbruck", "klagenfurt", "wels", "dornbirn"},
	},
	"finland": QueryPreset{
		include: []string{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku", "rovaniemi", "jyväskylä", "lahti", "kuopio", "pori", "lappeenranta", "vaasa"},
	},
	"sweden": QueryPreset{
		include: []string{"sweden", "sverige", "stockholm", "malmö", "uppsala", "göteborg", "gothenburg"},
	},
	"norway": QueryPreset{
		include: []string{"norway", "norge", "oslo", "bergen", "trondheim"},
	},
	"germany": QueryPreset{
		include: []string{"germany", "deutschland", "berlin", "frankfurt", "munich", "münchen", "hamburg", "cologne", "köln"},
	},
	"netherlands": QueryPreset{
		include: []string{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland", "delft"},
	},
	"ukraine": QueryPreset{
		include: []string{"ukraine", "kiev", "kyiv", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
	},
	"japan": QueryPreset{
		include: []string{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
	},
	"russia": QueryPreset{
		include: []string{"russia", "moscow", "saint+petersburg", "novosibirsk", "yekaterinburg", "nizhny+novgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
	},
	"estonia": QueryPreset{
		include: []string{"estonia", "eesti", "tallinn", "tartu", "narva", "pärnu"},
	},
	"denmark": QueryPreset{
		include: []string{"denmark", "danmark", "copenhagen", "aarhus", "odense", "aalborg"},
	},
	"portugal": QueryPreset{
		include: []string{"portugal", "lisbon", "lisboa", "braga", "porto", "aveiro", "coimbra", "funchal", "madeira"},
	},
	"france": QueryPreset{
		include: []string{"france", "paris", "marseille", "lyon", "toulouse", "nice", "nantes", "strasbourg", "montpellier", "bordeaux", "lille", "rennes", "reims", "rouen", "toulon", "le+havre", "grenoble", "dijon", "le+mans", "brest,france", "tours"},
	},
	"spain": QueryPreset{
		include: []string{"spain", "españa", "madrid", "barcelona", "valencia", "seville", "sevilla", "zaragoza", "malaga", "murcia", "palma", "bilbao", "alicante", "cordoba"},
	},
	"italy": QueryPreset{
		include: []string{"italy", "italia", "rome", "roma", "milan", "naples", "napoli", "turin", "torino", "palermo", "genoa", "genova", "bologna", "florence", "firenze", "bari", "catania", "venice", "verona"},
	},
	"uk": QueryPreset{
		include: []string{"uk", "england", "scotland", "wales", "northern+ireland", "london", "birmingham", "leeds", "glasgow", "sheffield", "bradford", "manchester", "edinburgh", "liverpool", "bristol", "cardiff", "belfast", "leicester", "wakefield", "coventry", "nottingham", "newcastle"},
	},
	"croatia": QueryPreset{
		include: []string{"croatia", "hrvatska", "zagreb", "split", "rijeka", "osijek", "zadar", "pula"},
	},
	"worldwide": QueryPreset{
		include: []string{},
	},
	"china": QueryPreset{
		include: []string{"china", "中国", "guangzhou", "shanghai", "beijing", "hangzhou", "hong+kong"},
	},
	"india": QueryPreset{
		include: []string{"india", "mumbai", "delhi", "bangalore", "hyderabad", "ahmedabad", "chennai", "kolkata", "jaipur"},
	},
	"israel": QueryPreset{
		include: []string{"israel", "tel+aviv", "jerusalem", "beer+sheva", "beersheva", "netanya", "ramat+gan", "haifa", "herzliya", "rishon"},
	},
	"indonesia": QueryPreset{
		include: []string{"indonesia", "jakarta", "surabaya", "bandung", "medan", "bekasi", "semarang", "tangerang", "depok", "makassar", "palembang"},
	},
	"pakistan": QueryPreset{
		include: []string{"pakistan", "karachi", "lahore", "faisalabad", "rawalpindi", "peshawar", "islamabad"},
	},
	"brazil": QueryPreset{
		include: []string{"brazil", "brasil", "são+paulo", "brasília", "salvador", "fortaleza", "belo+horizonte", "manaus", "curitiba", "recife", "porto+alegre", "florianópolis"},
	},
	"nigeria": QueryPreset{
		include: []string{"nigeria", "lagos", "kano", "ibadan", "benin+city", "port+harcourt", "jos", "ilorin"},
	},
	"bangladesh": QueryPreset{
		include: []string{"bangladesh", "dhaka", "chittagong", "khulna", "rajshahi", "barisal", "sylhet", "rangpur", "comilla", "gazipur"},
	},
	"mexico": QueryPreset{
		include: []string{"mexico", "mexico+city", "guadalajara", "puebla", "tijuana", "mexicali", "monterrey", "hermosillo", "zapopan", "ciudad+juarez", "chihuahua", "aguascalientes", "mx"},
	},
	"philippines": QueryPreset{
		include: []string{"philippines", "pilipinas", "quezon", "manila", "davao", "caloocan", "cebu", "zamboanga", "bohol", "pasig", "bacolod", "makati", "baguio", "cavite"},
	},
	"luxembourg": QueryPreset{
		include: []string{"luxembourg", "esch-sur-alzette", "differdange", "dudelange", "ettelbruck", "diekirch", "wiltz", "echternach", "rumelange", "grevenmacher", "bertrange", "mamer", "capellen", "strassen", "diekirch"},
	},
	"egypt": QueryPreset{
		include: []string{"egypt", "cairo", "alexandria", "giza", "port+said", "suez", "luxor", "el+mahalla", "asyut", "al+mansurah", "tanda"},
		exclude: []string{",+VA", "Virginia", ",+LA", "Louisiana"},
	},
	"ethiopia": QueryPreset{
		include: []string{"ethiopia", "addis+ababa", "gondar", "adama", "hawassa", "bahir+dar"},
	},
	"vietnam": QueryPreset{
		include: []string{"vietnam", "viet+nam", "ho+chi+minh", "hanoi", "ha+noi", "hai+phong", "da+nang", "can+tho", "bien+hoa", "nha+trang", "vinh"},
	},
	"iran": QueryPreset{
		include: []string{"iran", "tehran", "mashhad", "isfahan", "esfahan", "karaj", "shiraz", "tabriz", "qom", "ahvaz", "ahwaz", "kermanshah", "urmia", "rasht", "kerman"},
	},
	"congo": QueryPreset{
		include: []string{"congo", "drc", "kinshasa", "lubumbashi", "bukavu", "kananga", "goma"},
	},
	"turkey": QueryPreset{
		include: []string{"turkey", "turkiye", "istanbul", "ankara", "izmir", "bursa", "adana", "gaziantep", "konya", "antalya", "kayseri", "mersin", "eskisehir", "samsun", "denizli", "malatya"},
	},
	"thailand": QueryPreset{
		include: []string{"thailand", "bangkok", "nonthaburi", "nakhon", "phuket", "pattaya", "chiang+mai"},
	},
	"south africa": QueryPreset{
		include: []string{"south+africa", "south+africa", "johannesburg", "cape+town", "rsa", "durban", "port+elizabeth", "pretoria", "nelspruit"},
	},
	"myanmar": QueryPreset{
		include: []string{"myanmar", "burma", "yangon", "rangoon", "mandalay", "nay+pyi+taw", "taunggyi", "bago", "mawlamyine"},
	},
	"tanzania": QueryPreset{
		include: []string{"tanzania", "dar+es+salaam", "mwanza", "arusha", "dodoma", "mbeya", "morogoro", "tanga", "kilimanjaro"},
	},
	"south korea": QueryPreset{
		include: []string{"south+korea", "ROK", "korea", "seoul", "busan", "incheon", "daegu", "daejeon", "gwangju", "대한민국", "서울", "서울시"},
	},
	"colombia": QueryPreset{
		include: []string{"colombia", "bogota", "medellin", "cali", "barranquilla", "cartagena", "cucuta", "bucaramanga", "ibague", "soledad", "pereira", "santa+marta"},
	},
	"kenya": QueryPreset{
		include: []string{"kenya", "nairobi", "mombasa", "kisumu", "nakuru", "eldoret", "kisii", "nyeri"},
	},
	"argentina": QueryPreset{
		include: []string{"argentina", "buenos+aires", "cordoba", "rosario", "mendoza", "la+plata", "tucuman", "mar+del+plata", "salta", "resistencia"},
	},
	"algeria": QueryPreset{
		include: []string{"algeria", "algiers", "oran", "constantine", "annaba", "blida", "batna", "djelfa", "setif", "sidi+bel+abbes", "biskra", "tiaret", "relizane", "mostaganem", "tlemcen", "chlef", "jijel"},
	},
	"sudan": QueryPreset{
		include: []string{"sudan", "khartoum", "omdurman"},
	},
	"poland": QueryPreset{
		include: []string{"poland", "polska", "warsaw", "krakow", "lodz", "wroclaw", "poznan", "gdansk", "szczecin", "bydgoszcz", "lublin", "katowice", "bialystok"},
	},
	"canada": QueryPreset{
		include: []string{"canada", "ottawa", "edmonton", "winnipeg", "vancouver", "toronto", "quebec", "montreal", "mississauga", "calgary"},
	},
	"australia": QueryPreset{
		include: []string{"australia", "sydney", "melbourne", "brisbane", "perth", "adelaide", "canberra", "hobart"},
	},
	"new zealand": QueryPreset{
		include: []string{"new+zealand", "auckland", "wellington", "christchurch", "hamilton", "tauranga", "napier-hastings", "dunedin", "palmerston+north", "nelson", "rotorua", "whangarei", "new+plymouth", "invercargill", "whanganui", "gisborne"},
	},
	"belgium": QueryPreset{
		include: []string{"belgium", "antwerp", "ghent", "charleroi", "liege", "brussels", "belgique"},
	},
	"greece": QueryPreset{
		include: []string{"greece", "Ελλάδα", "athens", "thessaloniki", "patras", "heraklion", "larissa", "volos", "rhodes", "ioannina", "chania", "crete"},
		exclude: []string{"GA"},
	},
	"peru": QueryPreset{
		include: []string{"peru", "lima", "cusco", "cuzco", "ica", "arequipa", "trujillo", "chiclayo", "huancayo", "piura", "chimbote", "iquitos", "juliaca", "cajamarca"},
	},
	"hungary": QueryPreset{
		include: []string{"hungary", "magyarország", "budapest", "szeged", "miskolc"},
	},
	"albania": QueryPreset{
		include: []string{"albania", "tirana", "durres", "vlore", "elbasan", "shkoder"},
	},
	"uganda": QueryPreset{
		include: []string{"uganda", "kampala", "mbarara", "mukono", "jinja", "arua", "gulu", "masaka"},
	},
	"zambia": QueryPreset{
		include: []string{"zambia", "lusaka", "kitwe", "ndola"},
	},
	"sri lanka": QueryPreset{
		include: []string{"sri+lanka", "balangoda", "ratnapura", "colombo", "moratuwa", "negombo", "galle", "jaffna"},
	},
	"singapore": QueryPreset{
		include: []string{"singapore"},
	},
	"latvia": QueryPreset{
		include: []string{"latvia", "latvija", "riga", "rīga", "kuldiga", "kuldīga", "ventspils", "liepaja", "liepāja", "daugavpils", "jelgava", "jurmala", "jūrmala"},
	},
	"romania": QueryPreset{
		include: []string{"romania", "bucharest", "cluj", "iasi", "timisoara", "craiova", "brasov", "sibiu", "constanta", "oradea", "galati", "ploesti", "pitesti", "arad", "bacau"},
	},
	"belarus": QueryPreset{
		include: []string{"belarus", "minsk", "brest,belarus", "grodno", "gomel", "vitebsk", "mogilev", "slutsk", "borisov", "pinsk", "baranovichi", "bobruisk", "soligorsk"},
	},
	"malta": QueryPreset{
		include: []string{"malta", "birgu", "bormla", "mdina", "qormi", "senglea", "siġġiewi", "valletta", "zabbar", "zebbuġ", "zejtun"},
	},
	"rwanda": QueryPreset{
		include: []string{"rwanda", "kigali", "butare", "muhanga", "ruhengeri", "gisenyi"},
	},
	"saudi arabia": QueryPreset{
		include: []string{"Saudi", "KSA", "Riyadh", "Mecca"},
	},
	"morocco": QueryPreset{
		include: []string{"morocco", "casablanca", "fez", "tangier", "marrakesh", "salé", "meknes", "rabat", "oujda", "kenitra", "agadir", "tetouan", "temara", "safi", "mohammedia", "khouribga", "el+jadida"},
	},
	"uzbekistan": QueryPreset{
		include: []string{"uzbekistan", "tashkent", "namangan", "samarkand", "andijan", "nukus", "bukhara", "qarshi", "fergana"},
	},
	"malaysia": QueryPreset{
		include: []string{"malaysia", "kuala+lumpur", "kajang", "klang", "subang", "penang", "ipoh", "selangor", "melaka", "johor", "sabah"},
	},
	"afghanistan": QueryPreset{
		include: []string{"afghanistan", "kabul", "kandahar", "herat", "mazar-e-sharif", "jalalabad", "ghazni", "nangarhar"},
	},
	"venezuela": QueryPreset{
		include: []string{"venezuela", "caracas", "maracaibo", "barquisimeto", "guayana", "maturín", "zulia", "bolivar"},
	},
	"ghana": QueryPreset{
		include: []string{"ghana", "accra", "kumasi", "sekondi", "ashaiman", "sunyani", "tamale", "tema"},
	},
	"angola": QueryPreset{
		include: []string{"angola", "luanda", "huambo", "lobito", "benguela"},
	},
	"nepal": QueryPreset{
		include: []string{"nepal", "kathmandu", "pokhara", "lalitpur", "bharatpur", "birgunj", "biratnagar", "janakpur", "ghorahi"},
	},
	"yemen": QueryPreset{
		include: []string{"yemen", "sana'a", "taiz", "aden", "mukalla", "ibb"},
	},
	"mozambique": QueryPreset{
		include: []string{"mozambique", "maputo", "matola", "nampula", "beira", "sofala", "chimoio", "tete", "quelimane"},
	},
	"ivory coast": QueryPreset{
		include: []string{"ivory", "abidjan", "bouaké", "daloa", "yamoussoukro"},
	},
	"cameroon": QueryPreset{
		include: []string{"cameroon", "Douala", "Yaoundé", "Bafoussam", "Bamenda", "Garoua", "Maroua", "Ngaoundéré", "Kumba", "Nkongsamba", "Buea"},
	},
	"taiwan": QueryPreset{
		include: []string{"taiwan", "Taichung", "Kaohsiung", "Taipei", "Taoyuan", "Tainan", "Hsinchu", "Keelung", "Chiayi", "Changhua"},
	},
	"niger": QueryPreset{
		include: []string{"niger", "Niamey", "Maradi", "Zinder", "Tahoua", "Agadez", "Arlit", "Birni-N'Konni", "Dosso", "Gaya", "Tessaoua"},
	},
	"burkina faso": QueryPreset{
		include: []string{"burkina+faso", "Ouagadougou", "Bobo-Dioulasso", "Koudougou", "Banfora", "Ouahigouya", "Pouytenga", "Kaya", "Tenkodogo", "Fada+N'gourma", "Houndé"},
	},
	"mali": QueryPreset{
		include: []string{"mali", "bamako", "sikasso", "kalabancoro", "koutiala", "ségou", "kayes", "kati", "mopti", "niono"},
	},
	"malawi": QueryPreset{
		include: []string{"malawi", "Lilongwe", "Blantyre", "Mzuzu", "Zomba", "Karonga", "Kasungu", "Mangochi", "Salima", "Liwonde", "Balaka"},
	},
	"chile": QueryPreset{
		include: []string{"chile", "Santiago", "Valparaíso", "Concepción", "La+Serena", "Antofagasta", "Temuco", "Rancagua", "Talca", "Arica", "Chillán"},
	},
	"kazakhstan": QueryPreset{
		include: []string{"kazakhstan", "Almaty", "Shymkent", "Karagandy", "Taraz", "Nur-Sultan", "Pavlodar", "Oskemen", "Semey"},
	},
	"guatemala": QueryPreset{
		include: []string{"Guatemala", "mixco", "villa+nueva", "petapa", "Quetzaltenango"},
	},
	"ecuador": QueryPreset{
		include: []string{"ecuador", "Guayaquil", "Quito", "Cuenca", "Machala"},
	},
	"syria": QueryPreset{
		include: []string{"syria", "aleppo", "homs", "latakia", "hama", "raqqa"},
	},
	"cambodia": QueryPreset{
		include: []string{"cambodia", "phnom", "battambang", "siem+reap", "kampong"},
	},
	"senegal": QueryPreset{
		include: []string{"senegal", "dakar", "touba", "thies", "rufisque"},
	},
	"chad": QueryPreset{
		include: []string{"chad", "tchad", "n'djamena", "moundou"},
	},
	"somalia": QueryPreset{
		include: []string{"somalia", "mogadishu", "hargeisa", "bosaso", "borama"},
	},
	"zimbabwe": QueryPreset{
		include: []string{"zimbabwe", "harare", "bulawayo", "mutare", "gweru", "kwekwe"},
	},
	"guinea": QueryPreset{
		include: []string{"conakry"},
	},
	"benin": QueryPreset{
		include: []string{"benin", "cotonou", "porto-novo", "abomey"},
	},
	"haiti": QueryPreset{
		include: []string{"haiti", "port-au-prince", "cap-haitien", "carrefour", "delmas", "petion-ville"},
	},
	"cuba": QueryPreset{
		include: []string{"cuba", "havana", "santiago+de+cuba", "camaguey", "holguin", "guantanamo", "bayamo"},
	},
	"bolivia": QueryPreset{
		include: []string{"bolivia", "santa+cruz+de+la+sierra", "el+alto", "la+paz", "cochabamba", "oruro", "sucre"},
	},
	"tunisia": QueryPreset{
		include: []string{"tunisia", "tunis", "sfax", "sousse", "kairouan", "ariana", "gabes", "bizerte"},
	},
	"south sudan": QueryPreset{
		include: []string{"south sudan", "juba"},
	},
	"burundi": QueryPreset{
		include: []string{"burundi", "bujumbura", "gitega"},
	},
	"dominican republic": QueryPreset{
		include: []string{"dominican+republic", "republica+dominicana", "santo+domingo", "la+vega", "macoris"},
	},
	"czech republic": QueryPreset{
		include: []string{"czech", "ceska", "prague", "budejovice", "plzen", "karlovy", "ostrava", "brno"},
	},
	"jordan": QueryPreset{
		include: []string{"jordan", "amman", "zarqa", "irbid"},
	},
	"azerbaijan": QueryPreset{
		include: []string{"azerbaijan", "baku", "sumqayit", "ganja", "lankaran"},
	},
	"uae": QueryPreset{
		include: []string{"uae", "emirates", "dubai", "abu+dhabi", "sharjah", "al+ain", "ajman"},
	},
	"honduras": QueryPreset{
		include: []string{"honduras", "tegucigalpa", "san+pedro+sula", "choloma", "la+ceiba", "el+progreso", "choluteca", "comayagua"},
	},
	"tajikistan": QueryPreset{
		include: []string{"tajikistan", "dushanbe", "khujand"},
	},
	"papua new guinea": QueryPreset{
		include: []string{"papua+new+guinea", "port+moresby", "lae"},
	},
	"serbia": QueryPreset{
		include: []string{"serbia", "belgrade", "novi+sad", "nis", "kragujevac", "subotica", "zrenjanin", "pancevo", "cacak", "novi+pazar", "kraljevo", "smederevo"},
	},
	"switzerland": QueryPreset{
		include: []string{"switzerland", "zurich", "zürich", "geneva", "basel", "lausanne", "bern", "winterthur", "lucerne", "gallen", "lugano", "biel", "thun"},
	},
	"togo": QueryPreset{
		include: []string{"togo", "lome"},
	},
	"sierra leone": QueryPreset{
		include: []string{"sierra+leone", "freetown", "makeni", "koidu"},
	},
	"ireland": QueryPreset{
		include: []string{"ireland", "dublin", "cork", "limerick", "galway", "waterford+ireland", "drogheda", "dundalk"},
	},
	"hong kong": QueryPreset{
		include: []string{"hong+kong", "kowloon"},
	},
	"el salvador": QueryPreset{
		include: []string{"el+salvador"},
	},
	"kyrgyzstan": QueryPreset{
		include: []string{"kyrgyzstan", "bishkek", "osh", "jalal-abad", "karakol", "tokmok"},
	},
	"nicaragua": QueryPreset{
		include: []string{"nicaragua", "managua", "matagalpa", "chinandega"},
	},
	"turkmenistan": QueryPreset{
		include: []string{"turkmenistan", "turkmenabat"},
	},
	"paraguay": QueryPreset{
		include: []string{"paraguay", "asunción", "asuncion", "ciudad+del+este", "san+lorenzo", "luque", "capiata"},
	},
	"laos": QueryPreset{
		include: []string{"laos", "vientiane", "pakse"},
	},
	"bulgaria": QueryPreset{
		include: []string{"bulgaria", "sofia", "plovdiv", "varna", "burgas", "ruse", "stara+zagora", "pleven"},
	},
	"lebanon": QueryPreset{
		include: []string{"lebanon", "beirut", "sidon", "tyre"},
	},
	"libya": QueryPreset{
		include: []string{"libya", "tripoli", "benghazi", "misrata", "zliten", "bayda"},
	},
	"slovakia": QueryPreset{
		include: []string{"slovakia", "bratislava", "kosice", "presov", "zilina"},
	},
	"lithuania": QueryPreset{
		include: []string{"lithuania", "vilnius", "kaunas", "klaipeda", "siauliai", "panevezys", "alytus"},
	},
	"united states": QueryPreset{
		include: []string{",+US", "USA", "United+States", "Alabama", ",+AL", "Alaska", ",+AK", "Arizona", ",+AZ", "Arkansas", ",+AR", "California", ",+CA", "Colorado", ",+CO", "Connecticut", ",+CT", "Delaware", ",+DE", "Florida", ",+FL", "Georgia", ",+GA", "Hawaii", ",+HI", "Idaho", ",+ID", "Illinois", ",+IL", "Indiana", ",+IN", "Iowa", ",+IA", "Kansas", ",+KS", "Kentucky", ",+KY", "Louisiana", ",+LA", "Maine", ",+ME", "Maryland", ",+MD", "Massachusetts", ",+MA", "Michigan", ",+MI", "Minnesota", ",+MN", "Mississippi", ",+MS", "Missouri", ",+MO", "Montana", ",+MT", "Nebraska", ",+NE", "Nevada", ",+NV", "New+Hampshire", ",+NH", "New+Jersey", ",+NJ", "New+Mexico", ",+NM", "New+York", ",+NY", "North+Carolina", ",+NC", "North+Dakota", ",+ND", "Ohio", ",+OH", "Oklahoma", ",+OK", "Oregon", ",+OR", "Pennsylvania", ",+PA", "Rhode+Island", ",+RI", "South+Carolina", ",+SC", "South+Dakota", ",+SD", "Tennessee", ",+TN", "Texas", ",+TX", "Utah", ",+UT", "Vermont", ",+VT", "Virginia", ",+VA", "Washington", ",+WA", "West+Virginia", ",+WV", "Wisconsin", ",+WI", "Wyoming", ",+WY", "Los+Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia", "San+Antonio", "San+Diego", "Dallas", "San+Jose", "Austin", "Jacksonville", "Fort+Worth", "Columbus", "Charlotte", "San+Francisco", "Indianapolis", "Seattle", "Denver", "Boston", "El+Paso", "Nashville", "Detroit", "Portland", "Las+Vegas", "Memphis", "Louisville", "Baltimore"},
	},
}

func Preset(name string) QueryPreset {
	return PRESETS[name]
}
