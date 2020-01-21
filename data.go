package ausholiday

// csvStr holiday data from https://data.gov.au/dataset/ds-dga-b1bc6077-dadd-4f61-9f8c-002ab2cdff10/details
const csvStr = `Date,Holiday Name,Information,More Information,Jurisdiction
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.cmd.act.gov.au/communication/holidays,act
20190128,Australia Day,Always celebrated on 26 January,http://www.cmd.act.gov.au/communication/holidays,act
20190311,Canberra Day,"Held on the second Monday of March each year in Canberra, to celebrate the naming ceremony of AustraliaÅs capital which took place on 12 March 1913.",http://www.cmd.act.gov.au/communication/holidays,act
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.cmd.act.gov.au/communication/holidays,act
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.cmd.act.gov.au/communication/holidays,act
20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.cmd.act.gov.au/communication/holidays,act
20190527,Reconciliation Day,"take place on the first Monday on or after 27 May each year, the anniversary of the 1967 referendum",http://www.cmd.act.gov.au/communication/holidays,act
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.cmd.act.gov.au/communication/holidays,act
20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.cmd.act.gov.au/communication/holidays,act
20191225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.cmd.act.gov.au/communication/holidays,act
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.cmd.act.gov.au/communication/holidays,act
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190128,Australia Day,Always celebrated on 26 January,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190806,Bank Holiday,"Applies to banks and certain financial institutions, per the�Retail Trading Act 2008",http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20191225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190128,Australia Day,Always celebrated on 26 January,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190422,Easter Monday,Public Holiday as part of Easter.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190425,Anzac Day,Celebrated on the 25 April each year.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190506,May Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190805,Picnic Day,"Observed on the first Monday of August each year and is clebrated with a horse race, railway picnic and other social outings",https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20191224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20191231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190128,Australia Day,Always celebrated on 26 January,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190506,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20191007,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.safework.sa.gov.au/show_page.jsp?id=2483,sa
20190128,Australia Day,Always celebrated on 26 January,http://www.safework.sa.gov.au/show_page.jsp?id=2484,sa
20190311,March public holiday,The Holidays Act 1910 provides for the third Monday in May to be a public holiday. ch through the issuing of a special Proclamation by the Governor.,http://www.safework.sa.gov.au/show_page.jsp?id=2485,sa
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.safework.sa.gov.au/show_page.jsp?id=2486,sa
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.safework.sa.gov.au/show_page.jsp?id=2487,sa
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.safework.sa.gov.au/show_page.jsp?id=2488,sa
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.safework.sa.gov.au/show_page.jsp?id=2489,sa
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.safework.sa.gov.au/show_page.jsp?id=2490,sa
20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.safework.sa.gov.au/show_page.jsp?id=2491,sa
20191224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,http://www.safework.sa.gov.au/show_page.jsp?id=2492,sa
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.safework.sa.gov.au/show_page.jsp?id=2493,sa
20191226,Proclamation Day ,Occurs the day after Christmas Day.,http://www.safework.sa.gov.au/show_page.jsp?id=2494,sa
20191231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,http://www.safework.sa.gov.au/show_page.jsp?id=2495,sa
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190128,Australia Day,Always celebrated on 26 January,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190211,Royal Hobart Regatta,A three-day event that runs over the weekend of the second Monday each February.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190311,Eight Hours Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://worksafe.tas.gov.au/laws/public_holidays,tas
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190422,Easter Monday,Public Holiday as part of Easter.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190423,Easter Tuesday,Public Holiday currently observed by certain awards/agreements and the State Public Service,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190425,Anzac Day,Celebrated on the 25 April each year.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://worksafe.tas.gov.au/laws/public_holidays,tas
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190128,Australia Day,Always celebrated on 26 January,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190311,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190927,Friday before AFL Grand Final,Friday before the AFL Grand Final,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20191105,Melbourne Cup,All of Victoria unless alternate local holiday has been arranged by non-metro council.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190128,Australia Day,Always celebrated on 26 January,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190304,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190603,Western Australia Day,Held on the first Monday in June each year and is a state holiday only.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20190930,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.cmtedd.act.gov.au/communication/holidays,act
20200127,Australia Day,Always celebrated on 26 January,https://www.cmtedd.act.gov.au/communication/holidays,act
20200309,Canberra Day,"Held on the second Monday of March each year in Canberra, to celebrate the naming ceremony of AustraliaÃÂs capital which took place on 12 March 1913.",https://www.cmtedd.act.gov.au/communication/holidays,act
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200411,Easter Saturday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200412,Easter Sunday,Public Holiday as part of Easter.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200601,Reconciliation Day,"take place on the first Monday on or after 27 May each year, the anniversary of the 1967 referendum",https://www.cmtedd.act.gov.au/communication/holidays,act
20200608,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.cmtedd.act.gov.au/communication/holidays,act
20201005,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.cmtedd.act.gov.au/communication/holidays,act
20201225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.cmtedd.act.gov.au/communication/holidays,act
20201228,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.cmtedd.act.gov.au/communication/holidays,act
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200127,Australia Day,Always celebrated on 26 January,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200411,Easter Saturday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200412,Easter Sunday,Public Holiday as part of Easter.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200803,Bank Holiday,"Applies to banks and certain financial institutions, per theÊRetail Trading Act 2008",https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20201005,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20201225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20201226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20201228,Boxing Day (Additional day),Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.industrialrelations.nsw.gov.au/public-holidays/public-holidays-nsw,nsw
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://nt.gov.au/nt-public-holidays,nt
20200127,Australia Day,Always celebrated on 26 January,https://nt.gov.au/nt-public-holidays,nt
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://nt.gov.au/nt-public-holidays,nt
20200411,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://nt.gov.au/nt-public-holidays,nt
20200413,Easter Monday,Public Holiday as part of Easter.,https://nt.gov.au/nt-public-holidays,nt
20200425,Anzac Day,Celebrated on the 25 April each year.,https://nt.gov.au/nt-public-holidays,nt
20200504,May Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://nt.gov.au/nt-public-holidays,nt
20200608,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://nt.gov.au/nt-public-holidays,nt
20200803,Picnic Day,"Observed on the first Monday of August each year and is clebrated with a horse race, railway picnic and other social outings",https://nt.gov.au/nt-public-holidays,nt
20201224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/nt-public-holidays,nt
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://nt.gov.au/nt-public-holidays,nt
20201228,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://nt.gov.au/nt-public-holidays,nt
20201231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/nt-public-holidays,nt
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200127,Australia Day,Always celebrated on 26 January,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200411,Easter Saturday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200412,Easter Sunday,Public Holiday as part of Easter.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200504,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.qld.gov.au/recreation/travel/holidays/public,qld
20201005,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20201226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20201228,Boxing Day (Additional day),Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.qld.gov.au/recreation/travel/holidays/public,qld
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200126,Australia Day,Always celebrated on 26 January,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200127,Australia Day (Additional day),Always celebrated on 26 January,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200309,Adelaide Cup Day,The Holidays Act 1910 provides for the third Monday in May to be a public holiday. ch through the issuing of a special Proclamation by the Governor.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200411,Easter Saturday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200608,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20201005,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20201224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20201228,Proclamation Day ,Occurs the day after Christmas Day.,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20201231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,https://www.safework.sa.gov.au/law-compliance/laws-regulations/public-holidays,sa
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200127,Australia Day,Always celebrated on 26 January,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200309,Eight Hours Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://worksafe.tas.gov.au/laws/public_holidays,tas
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200413,Easter Monday,Public Holiday as part of Easter.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200414,Easter Tuesday,Public Holiday currently observed by certain awards/agreements and the State Public Service,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200425,Anzac Day,Celebrated on the 25 April each year.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200608,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20201228,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://worksafe.tas.gov.au/laws/public_holidays,tas
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2020,vic
20200127,Australia Day,Always celebrated on 26 January,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2021,vic
20200309,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2022,vic
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2023,vic
20200411,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2024,vic
20200412,Easter Sunday,Public Holiday as part of Easter.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2025,vic
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2026,vic
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2027,vic
20200608,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2028,vic
20200925,Friday before AFL Grand Final,Friday before the AFL Grand Final,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2029,vic
20201103,Melbourne Cup,All of Victoria unless alternate local holiday has been arranged by non-metro council.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2030,vic
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2031,vic
20201226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2032,vic
20201228,Boxing Day (Additional day),Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays-2032,vic
20200101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200127,Australia Day,Always celebrated on 26 January,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200302,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200410,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200413,Easter Monday,Public Holiday as part of Easter.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200425,Anzac Day,Celebrated on the 25 April each year.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200427,Anzac Day,Celebrated on the 25 April each year.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200601,Western Australia Day,Held on the first Monday in June each year and is a state holiday only.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20200928,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20201225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20201226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
20201228,Boxing Day (Additional day),Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
`

// csvStrTest data for testing
const csvStrTest = `Date,Holiday Name,Information,More Information,Jurisdiction
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.cmd.act.gov.au/communication/holidays,act
20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.cmd.act.gov.au/communication/holidays,act
20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
20190422,Easter Monday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.cmd.act.gov.au/communication/holidays,act
20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
`
