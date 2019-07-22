package ausholiday

// csvStr holiday data from https://data.gov.au/dataset/ds-dga-b1bc6077-dadd-4f61-9f8c-002ab2cdff10/details
const csvStr = `Raw Date,Date,Holiday Name,Information,More Information,Jurisdiction
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.cmd.act.gov.au/communication/holidays,act
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.cmd.act.gov.au/communication/holidays,act
1552222800,20190311,Canberra Day,"Held on the second Monday of March each year in Canberra, to celebrate the naming ceremony of AustraliaÅs capital which took place on 12 March 1913.",http://www.cmd.act.gov.au/communication/holidays,act
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.cmd.act.gov.au/communication/holidays,act
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.cmd.act.gov.au/communication/holidays,act
1555768800,20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.cmd.act.gov.au/communication/holidays,act
1558879200,20190527,Reconciliation Day,"take place on the first Monday on or after 27 May each year, the anniversary of the 1967 referendum",http://www.cmd.act.gov.au/communication/holidays,act
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.cmd.act.gov.au/communication/holidays,act
1570366800,20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.cmd.act.gov.au/communication/holidays,act
1577192400,20191225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.cmd.act.gov.au/communication/holidays,act
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.cmd.act.gov.au/communication/holidays,act
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1555768800,20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1565013600,20190806,Bank Holiday,"Applies to banks and certain financial institutions, per the�Retail Trading Act 2008",http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1570366800,20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1577192400,20191225,Christmas Day ,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.industrialrelations.nsw.gov.au/oirwww/NSW_public_holidays/NSW_Public_Holidays.page?,nsw
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1548594000,20190128,Australia Day,Always celebrated on 26 January,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1557064800,20190506,May Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1564927200,20190805,Picnic Day,"Observed on the first Monday of August each year and is clebrated with a horse race, railway picnic and other social outings",https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1577106000,20191224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1577710800,20191231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,https://nt.gov.au/employ/for-employees-in-nt/nt-public-holidays,nt
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1555768800,20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1557064800,20190506,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1570366800,20191007,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.justice.qld.gov.au/fair-and-safe-work/industrial-relations/public-holidays/dates,qld
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.safework.sa.gov.au/show_page.jsp?id=2483,sa
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.safework.sa.gov.au/show_page.jsp?id=2484,sa
1552222800,20190311,March public holiday,The Holidays Act 1910 provides for the third Monday in May to be a public holiday. ch through the issuing of a special Proclamation by the Governor.,http://www.safework.sa.gov.au/show_page.jsp?id=2485,sa
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.safework.sa.gov.au/show_page.jsp?id=2486,sa
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.safework.sa.gov.au/show_page.jsp?id=2487,sa
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.safework.sa.gov.au/show_page.jsp?id=2488,sa
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.safework.sa.gov.au/show_page.jsp?id=2489,sa
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.safework.sa.gov.au/show_page.jsp?id=2490,sa
1570366800,20191007,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.safework.sa.gov.au/show_page.jsp?id=2491,sa
1577106000,20191224,Christmas Eve,Part-day public holiday from 7pm-12 midnight,http://www.safework.sa.gov.au/show_page.jsp?id=2492,sa
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.safework.sa.gov.au/show_page.jsp?id=2493,sa
1577278800,20191226,Proclamation Day ,Occurs the day after Christmas Day.,http://www.safework.sa.gov.au/show_page.jsp?id=2494,sa
1577710800,20191231,New Year's Eve,Part-day public holiday from 7pm-12 midnight,http://www.safework.sa.gov.au/show_page.jsp?id=2495,sa
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://worksafe.tas.gov.au/laws/public_holidays,tas
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://worksafe.tas.gov.au/laws/public_holidays,tas
1549803600,20190211,Royal Hobart Regatta,A three-day event that runs over the weekend of the second Monday each February.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1552222800,20190311,Eight Hours Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://worksafe.tas.gov.au/laws/public_holidays,tas
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1555941600,20190423,Easter Tuesday,Public Holiday currently observed by certain awards/agreements and the State Public Service,http://worksafe.tas.gov.au/laws/public_holidays,tas
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://worksafe.tas.gov.au/laws/public_holidays,tas
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1552222800,20190311,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1555768800,20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1560088800,20190610,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1569506400,20190927,Friday before AFL Grand Final,Friday before the AFL Grand Final,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1572872400,20191105,Melbourne Cup,All of Victoria unless alternate local holiday has been arranged by non-metro council.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.business.vic.gov.au/victorian-public-holidays-and-daylight-saving/victorian-public-holidays,vic
1546261200,20190101,New Year's Day,New Year's Day is the first day of the calendar year and is celebrated each January 1st,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1548594000,20190128,Australia Day,Always celebrated on 26 January,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1551618000,20190304,Labour Day,"Always on a Monday, creating a long weekend. It celebrates the eight-hour working day, a victory for workers in the mid-late 19th century.",http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1559484000,20190603,Western Australia Day,Held on the first Monday in June each year and is a state holiday only.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1569765600,20190930,Queen's Birthday,Celebrated on second Monday in June except in Western Australia and Queensland.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1577192400,20191225,Christmas Day,Christmas Day is an annual holiday which celebrates the birth of Jesus Christ over 2000 years ago.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
1577278800,20191226,Boxing Day,Boxing Day occurs the day after Christmas. Sydney-to-Hobart yacht race and Boxing Day Test Match (Cricket) start on this day.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
`

// csvStrTest data for testing
const csvStrTest = `Raw Date,Date,Holiday Name,Information,More Information,Jurisdiction
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.cmd.act.gov.au/communication/holidays,act
1555682400,20190420,Saturday before Easter Sunday,Easter Saturday is between Good Friday and Easter Sunday in Australia.,http://www.cmd.act.gov.au/communication/holidays,act
1555768800,20190421,Easter Sunday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
1555855200,20190422,Easter Monday,Public Holiday as part of Easter.,http://www.cmd.act.gov.au/communication/holidays,act
1556114400,20190425,Anzac Day,Celebrated on the 25 April each year.,http://www.cmd.act.gov.au/communication/holidays,act
1555596000,20190419,Good Friday,Easter is celebrated with Good Friday and Easter Monday creating a 4 day long weekend.,http://www.commerce.wa.gov.au/labour-relations/public-holidays-western-australia,wa
`
