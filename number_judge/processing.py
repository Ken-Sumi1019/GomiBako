import keras
import numpy as np
#予測モデルをインポート
model = keras.models.load_model('num-judge.h5')

#予測する関数
def predict(data):
    global model
    ans = model.predict(data)
    return ans[0]

#受け取った二次元配列を予測に使えるように
#四次元配列に変換する関数
def convert(ls):
    data = np.array([np.array([np.array([0]) for j in range(len(ls[0]))]) for i in range(len(ls))])
    for i in range(len(ls)):
        for j in range(len(ls[0])):
            data[i][j][0] = ls[i][j]
    return np.array([data])