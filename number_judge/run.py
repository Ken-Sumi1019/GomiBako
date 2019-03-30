from flask import *
import numpy as np
from processing import *
import numpy as np
app = Flask(__name__)

#メインのページを表示
@app.route('/')
def index():
    return render_template('mainpage.html')

#受け取った文字列をモデルに合うように整形して予測
#結果を表示するページにとぶ
@app.route('/predict',methods=["POST"])
def get_data():
    print('#')
    string = request.form['data']
    data = np.array([np.array([0 for i in range(28)]) for j in range(28)])
    for i in range(28*28):
        data[i//28][i%28] = string[i]
    data = convert(data)
    ans = predict(data)
    del data
    ok = 0;ma = 0
    for i in range(len(ans)):
        if ans[i] > ma:ok = i;ma = ans[i]
    print(ma,ok)
    return render_template('submit.html', prob = str(ma*100),ans = str(ok))

if __name__ == '__main__':
    app.run()