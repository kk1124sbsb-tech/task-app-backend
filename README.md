# Task Management App (Backend)

重要度と期限に基づいた、ストレスのないタスク管理を実現するためのバックエンドAPI。

## 開発の背景
アイゼンハワーマトリクスからインスピレーションを受け、「緊急度」だけでなく「重要度」という指標を導入し、今本当にやるべきことを整理したいというニーズから開発。
個人的に大学の課題やTOEICの勉強など、日々様々なタスクがあり、普段からアイゼンハワーマトリクスをアナログで用いてタスク管理することが多かったが、アナログだとマトリクスが汚れてきて、結局タスクが簡潔に管理できないという状況に陥っていたため、より管理しやすいようにアプリに落とし込こみたいと考えた。

## 主な機能（予定含む）
- タスクの重要度設定（3段階: 高・中・低）
- 期限（Deadline）によるソート機能
- 重要度によるフィルタリング

- タスクの重さも組み込みたい→保留

## 技術スタック
- **Language:** Go 1.22.2
- **Infrastructure:** Docker / PostgreSQL
- **Architecture:** REST API

## 実行方法
```bash
git clone https://github.com/kk1124sbsb-tech/task-app-backend
cd task-app-backend
go run main.go