/**
 * Task の型定義
 */
export interface Task {
    id: number;
    title: string;
    done: boolean;
}

/**
 * TaskManager クラス
 * カプセル化と配列メソッドの活用
 */
export class TaskManager {
    private tasks: Task[] = [];
    private nextId: number = 1;

    /**
     * タスク追加
     */
    add(title: string): void {
        const newTask: Task = {
            id: this.nextId++,
            title,
            done: false
        };
        this.tasks.push(newTask);
    }

    /**
     * タスク完了
     * .find() を使用したスマートな検索
     */
    complete(id: number): void {
        const task = this.tasks.find(t => t.id === id);
        if (!task) {
            throw new Error(`Task with id ${id} not found`);
        }
        task.done = true;
    }

    /**
     * タスク一覧
     * .map() と .forEach() を使用した表示
     */
    list(): void {
        if (this.tasks.length === 0) {
            console.log("No tasks available.");
            return;
        }

        this.tasks.forEach(t => {
            const status = t.done ? "[x]" : "[ ]";
            console.log(`${t.id}: ${status} ${t.title}`);
        });
    }

    /**
     * 完了済みのみを抽出 (filterのデモ用)
     */
    getCompletedTasks(): Task[] {
        return this.tasks.filter(t => t.done);
    }
}
