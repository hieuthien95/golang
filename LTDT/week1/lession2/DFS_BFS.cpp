// lab2.cpp : Defines the entry point for the console application.
//

#include "iostream"
#include "fstream"
#include "vector"

using namespace std;

#define max 100

class Graph
{
public:
	bool isVisited[max];

	// A: ma tran ke cua G, n: so dinh
	int A[max][max];
	int n;

	void ReadFile(char *fileName);
	void DFS_DeQuy(int u);
	void DFS(int u);
	void BFS(int u);
};

// doc file chua do thi G luu vao ma tran A
void Graph::ReadFile(char *fileName) {
	fstream f(fileName, ios::in);
	if (!f)
	{
		cout << "khong doc duoc file";
	}
	else
	{
		f >> n;
		cout << n << endl;
		for (int i = 0; i < n; i++) {
			for (int j = 0; j < n; j++) {
				f >> A[i][j];
				cout << A[i][j] << " ";
			}
			cout << endl;
		}
	}
	f.close();

	// danh top dinh chua xet
	for (int i = 0; i < max; i++)
		isVisited[i] = false;
}

// thuat toan DFS de quy
void Graph::DFS_DeQuy(int u){
	// xet dinh u
	isVisited[u] = true;
	cout << u + 1 << " ";
	for (int v = 0; v < n; v++)
	if (isVisited[v] == false && A[u][v] == 1)
	{
		DFS_DeQuy(v);
	}
}

// thuat toan DFS
void Graph::DFS(int s){
	// danh top dinh chua xet
	for (int i = 0; i < n; i++)
		isVisited[i] = false;

	int stack[max], top = 0;

	stack[top++] = s;
	// cout << stack;


	while (!top == 0) {
		int v = stack[--top];
		if (isVisited[v] == false) {
			// cout << v + 1 << " ";
			isVisited[v] = true;
			for (int i = n; i >= 1; i--) {
				if (isVisited[i] == false && A[v][i] != 0) {
					stack[top++] = i;
				}
			}
		}

	}
}

// thuat toan BFS
void Graph::BFS(int u){
	// danh top dinh chua xet
	for (int i = 0; i < n; i++)
		isVisited[i] = false;

	int queue[max], top = 0, bottom = 0;
	for (int i = 0; i < n; i++)
		queue[i] = 0;

	queue[bottom] = u;
	isVisited[u] = true;
	cout << u + 1 << " ";

	while (top >= bottom)
	{
		int p = queue[bottom];
		bottom++;
		for (int v = 0; v < n; v++)
		if (isVisited[v] == false && A[p][v] == 1)
		{
			top++;
			queue[top] = v;
			isVisited[v] = true;
			cout << v + 1 << " ";
		}
	}
}

int main() {
	Graph g = Graph();
	g.ReadFile("input.txt");

	cout << "\nDuyet do thi DFS de quy: ";
	g.DFS_DeQuy(0);

	cout << "\nDuyet do thi DFS: ";
	g.DFS(0);

	cout << "\nDuyet do thi BFS: ";
	g.BFS(0);

	system("pause");

	return 0;
}