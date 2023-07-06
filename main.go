package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var totalTime float64

func requestTest(serverAddress, authKey string) {
	start := time.Now()

	// url := serverAddress+"/directions/json?key="+authKey+"&origin=25.036413,121.526674&destination=25.049216,121.466643&is_send_detail_points=true&waypoints=25.035032,121.547499%7C25.043315,121.561721%7C25.055279,121.556133%7C25.055049,121.544197%7C25.05689,121.534293%7C25.065631,121.525404%7C25.067472,121.498993%7C25.066782,121.494675%7C25.070462,121.483247%7C25.085413,121.508389%7C25.086333,121.520325%7C25.042625,121.539372%7C25.048837,121.552578%7C25.06011,121.579498%7C25.074143,121.591434%7C25.080093,121.571503%7C25.081178,121.559224%7C25.08172,121.551438%7C25.100434,121.522389%7C25.114264,121.527181%7C25.11345,121.518796%7C25.051609,121.466987%7C25.065988,121.485853%7C25.056493,121.488549%7C25.019592,121.472377%7C25.004123,121.472976%7C24.996252,121.489747%7C24.999237,121.504121%7C25.002121,121.514021%7C24.990905,121.517203%7C25.026152,121.55256%7C25.029355,121.529932%7C25.028074,121.510839%7C25.028394,121.499525%7C25.059146,121.514375%7C25.071957,121.588978%7C25.043451,121.576603%7C25.035122,121.5713"
	// method := "GET"

	// payload := strings.NewReader("")

	url := serverAddress + "/sequence/json?key=" + authKey
	method := "POST"

	data := `{"origin":{"lat":24.9101262,"lng":121.2411366},"waypoints":[{"lat":24.9988141,"lng":121.159045},{"lat":25.0488341,"lng":121.1360643},{"lat":25.0209197,"lng":121.1296829},{"lat":24.9927111,"lng":121.1412987},{"lat":25.0114908,"lng":121.1603087},{"lat":25.043595,"lng":121.1356704},{"lat":24.9850764,"lng":121.1349377},{"lat":25.0298937,"lng":121.0593087},{"lat":25.05071,"lng":121.1181555},{"lat":25.0189221,"lng":121.0932281},{"lat":25.0410281,"lng":121.1211607},{"lat":25.0243082,"lng":121.1527118},{"lat":24.9883192,"lng":121.1617097},{"lat":25.047685,"lng":121.1302489},{"lat":25.036586,"lng":121.1383912},{"lat":25.045108,"lng":121.075452},{"lat":25.0437039,"lng":121.1355478},{"lat":24.9754337,"lng":121.1470256},{"lat":25.0520178,"lng":121.1225136},{"lat":25.0685825,"lng":121.1239596},{"lat":25.0233135,"lng":121.1338071},{"lat":25.0541264,"lng":121.1240367},{"lat":25.0463909,"lng":121.1356634},{"lat":25.0248078,"lng":121.0545095},{"lat":25.0426372,"lng":121.1369146},{"lat":25.0386061,"lng":121.1358544},{"lat":25.0448166,"lng":121.1577379},{"lat":25.040121,"lng":121.1605327},{"lat":25.034954,"lng":121.081261},{"lat":25.0122589,"lng":121.0898653},{"lat":25.0478114,"lng":121.084089},{"lat":25.0304947,"lng":121.1050659},{"lat":25.0360833,"lng":121.103446},{"lat":25.0521089,"lng":121.11805},{"lat":25.0353997,"lng":121.0828183},{"lat":25.0241023,"lng":121.1304561},{"lat":25.0122589,"lng":121.0898653},{"lat":25.0331067,"lng":121.0952133},{"lat":25.0547326,"lng":121.1239597},{"lat":25.0159304,"lng":121.1348306},{"lat":25.0593082,"lng":121.1271818},{"lat":25.0578377,"lng":121.120245},{"lat":25.0421085,"lng":121.1411962},{"lat":25.0298937,"lng":121.0593087},{"lat":25.0572868,"lng":121.1247249},{"lat":25.0070354,"lng":121.1123118},{"lat":25.0026839,"lng":121.1357142},{"lat":25.0432102,"lng":121.1381128},{"lat":25.042242,"lng":121.1410928},{"lat":25.0509379,"lng":121.1180157},{"lat":25.049253,"lng":121.1281452},{"lat":25.0110541,"lng":121.1392951},{"lat":25.0324265,"lng":121.1112697},{"lat":25.0281681,"lng":121.083837},{"lat":25.0131065,"lng":121.1359956},{"lat":25.0157563,"lng":121.1446038},{"lat":25.0302928,"lng":121.0742455},{"lat":24.9750176,"lng":121.1580547},{"lat":25.0518997,"lng":121.1185417},{"lat":24.981779,"lng":121.1361203},{"lat":25.0526174,"lng":121.1209858},{"lat":25.0336718,"lng":121.0668085},{"lat":25.0547326,"lng":121.1239597},{"lat":25.0350829,"lng":121.0808649},{"lat":25.0464907,"lng":121.1421876},{"lat":25.0421579,"lng":121.1432007},{"lat":25.0629338,"lng":121.1322555},{"lat":25.0531398,"lng":121.1473134},{"lat":25.0459016,"lng":121.1391303},{"lat":25.0359365,"lng":121.1137544},{"lat":25.0359365,"lng":121.1137544},{"lat":25.0154918,"lng":121.134373},{"lat":25.0117248,"lng":121.1051741},{"lat":24.978024,"lng":121.1616292},{"lat":25.0527646,"lng":121.1064359},{"lat":25.0440748,"lng":121.1385385},{"lat":25.014855,"lng":121.1320595},{"lat":25.0082399,"lng":121.0694853},{"lat":24.9745908,"lng":121.1504453},{"lat":25.053123,"lng":121.1469835},{"lat":24.9754337,"lng":121.1470256},{"lat":25.0462223,"lng":121.1351274},{"lat":25.015768,"lng":121.1561123},{"lat":24.9738627,"lng":121.1386682},{"lat":25.012157,"lng":121.1360833},{"lat":25.0488884,"lng":121.1376003},{"lat":25.0259315,"lng":121.1520715},{"lat":24.9975618,"lng":121.1181096},{"lat":25.0359365,"lng":121.1137544},{"lat":25.008756,"lng":121.142055},{"lat":24.9956084,"lng":121.1636824},{"lat":25.0525332,"lng":121.1244336},{"lat":25.053248,"lng":121.105655},{"lat":25.0246909,"lng":121.0555453},{"lat":25.0547242,"lng":121.1434874},{"lat":25.0032092,"lng":121.1023141},{"lat":25.0488869,"lng":121.1289521},{"lat":25.0135489,"lng":121.1249224}],"destination":{"lat":24.9101262,"lng":121.2411366},"limit_time_parameter":{"start_time":1682409654,"limit_time_points":[{"point_index":0,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":1,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":2,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":3,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":4,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":5,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":6,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":7,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":8,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":9,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":10,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":11,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":12,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":13,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":14,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":15,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":16,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":17,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":18,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":19,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":20,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":21,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":22,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":23,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":24,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":25,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":26,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":27,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":28,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":29,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":30,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":31,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":32,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":33,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":34,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":35,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":36,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":37,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":38,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":39,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":40,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":41,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":42,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":43,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":44,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":45,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":46,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":47,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":48,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":49,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":50,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":51,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":52,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":53,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":54,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":55,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":56,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":57,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":58,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":59,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":60,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":61,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":62,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":63,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":64,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":65,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":66,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":67,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":68,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":69,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":70,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":71,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":72,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":73,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":74,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":75,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":76,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":77,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":78,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":79,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":80,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":81,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":82,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":83,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":84,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":85,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":86,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":87,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":88,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":89,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":90,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":91,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":92,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":93,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":94,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":95,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":96,"limit_after_time":1682418600,"limit_before_time":1682438340},{"point_index":97,"limit_after_time":1682418600,"limit_before_time":1682438340}]}}`

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resJson map[string]interface{}
	err = json.Unmarshal(body, &resJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(resJson["status"])
	fmt.Println(res.StatusCode)
	fmt.Println(resJson["error_message"])
	elapsed := time.Since(start)
	fmt.Printf("success: %s\n", elapsed)
	totalTime += elapsed.Seconds()
}

func main() {
	var authKey = ""
	var serverAddress = ""
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestTest(serverAddress, authKey)
		}()
	}
	wg.Wait()
	fmt.Println("All requests finished", totalTime)
}
