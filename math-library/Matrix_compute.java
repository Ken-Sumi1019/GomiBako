public class Matrixcompute {

	// 行列の内積
	public static double[][] matrixdot(double a[][], double b[][]) {
		double ans[][] = new double[a.length][b[0].length];
		if (a.length != b[0].length) {
			throw new IllegalArgumentException("行列の積が計算できません");
		}

		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < b[0].length; j++) {
				double sum = 0;
				for (int k = 0; k < a[0].length; k++) {
					sum += a[i][k] * b[j][k];
				}
				ans[i][j] = sum;
			}
		}
		return ans;
	}

	// 転置行列
	public static double[][] matrixT(double a[][]) {
		double ans[][] = new double[a[0].length][a.length];

		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a[0].length; j++) {
				ans[j][i] = a[i][j];
			}
		}
		return ans;
	}

	// 行列に定数倍
	public static double[][] matrixmultiple(double a[][], double num) {
		double ans[][] = new double[a.length][a[0].length];

		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a[0].length; j++) {
				ans[i][j] = (double)a[i][j] * num;
			}
		}
		return ans;
	}

	// 行列の加算
	public static double[][] matrixsum(double a[][], double b[][]) {
		double ans[][] = new double[a.length][a[0].length];
		if (a.length != b.length || a[0].length != b[0].length) {
			throw new IllegalArgumentException("行列の形が違います");
		}
		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a[0].length; j++) {
				ans[i][j] = a[i][j] + b[i][j];
			}
		}
		return ans;
	}

	// 行列の減算
	public static double[][] matrixsub(double a[][], double b[][]) {
		double ans[][] = new double[a.length][a[0].length];
		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a[0].length; j++) {
				ans[i][j] = a[i][j] - b[i][j];
			}
		}
		return ans;
	}

	// 行列式を出す
	public static double matrixdet(double a[][]) {
		if (a.length != a[0].length) {
			throw new IllegalArgumentException("正方行列じゃないお");
		}
		if (a.length == 2) {
			double ans = a[0][0] * a[1][1] - a[0][1] * a[1][0];
			return ans;
		} else {
			double ans = 0;
			for (int i = 0; i < a[0].length; i++) {
				double save[][] = new double[a.length - 1][a.length - 1];
				int y_count = 0;
				for (int y = 0; y < a.length; y++) {
				    if(0 == y)continue;
				    int x_count = 0;
				    for (int x = 0;x < a.length;x++){
				        if(x==i)continue;
				        save[y_count][x_count] = a[y][x];
				        x_count++;
				    }
				    y_count++;

				}
				int num = 0;
				if (i % 2 == 0)
					num = 1;
				else
					num = -1;
				ans += a[0][i] * num * Matrixcompute.matrixdet(save);
			}
			return ans;
		}
	}

	// 逆行列の導出
	public static double[][] matrixinv(double a[][]) {
		double ans[][] = new double[a.length][a[0].length];
		if (a.length != a[0].length) {
			throw new IllegalArgumentException("正方行列じゃないお");
		}
		double det = Matrixcompute.matrixdet(a);
		System.out.println(det);
		for (int i = 0; i < a.length; i++) {
			for (int j = 0; j < a.length; j++) {
				double save[][] = new double[a.length - 1][a.length - 1];
				int count_x = 0, count_y = 0;
				for (int y = 0; y < a.length; y++) {
					if (y == i){
						continue;
					}else{
					    count_x = 0;
    					for (int x = 0; x < a.length; x++) {
	    					if (x == j){
			    				continue;
				    		}else{
					    	    save[count_y][count_x] = a[y][x];
						        count_x++;
						    }
    					}
					}
					count_y++;
				}
				int num = 0;
				if ((i + j) % 2 == 0)
					num = 1;
				else
					num = -1;
				ans[i][j] = num * Matrixcompute.matrixdet(save);
			}
		}
		ans = Matrixcompute.matrixT(ans);
		ans = Matrixcompute.matrixmultiple(ans,(double)1/det);
		return ans;
	}
}

