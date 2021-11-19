import "dart:io";

void main() {
    Program().run();
}

class Program {
    bool exit = false;
    bool isDeleting = false;
    Command cmd = Command("");
    String message = "";

    List<ScheduleItem> list = [
        ScheduleItem(331, "Hello, World!", "How are you?"),
        ScheduleItem(201, "Second Title", "")
    ];

    void run() {
        while (!exit) {
            draw();
            update();
        }
    }

    void draw() {
        UI.ClearScreen();

        UI.Text("┏━━ Study Schedule ━━━\n");
        UI.EmptyLine();

        if (list.isNotEmpty) {
            UI.Text("  List:\n");
            UI.ListView(list.length, (int i) {
                var item = list.elementAt(i);

                if (isDeleting) {
                    UI.Text("    [${i + 1}] ${item.title}\n");
                } else {
                    UI.Text("    ${item.title}\n");
                }

                if (item.desc.isNotEmpty) {
                    UI.Text("      ${item.desc}\n");
                    UI.EmptyLine();
                } else {
                    UI.EmptyLine();
                }
            });
        }

        UI.EmptyLine();
        UI.EmptyLine();

        if (message.isNotEmpty) {
            UI.Text(message);
        } else {
            UI.EmptyLine();
        }

        UI.Text(">> ");
    }

    void update() {
        message = "";

        cmd = Command(stdin.readLineSync() ?? "");
        if (cmd.name.isEmpty) return;

        switch (cmd.name) {
            case "add":
            case "ad":
            case "a": {
                draw();
                stdout.write("Title (empty to cancel): ");
                var title = stdin.readLineSync() ?? "";
                title = title.trim();

                if (title.isEmpty) {
                    message = "Cancelled\n";
                    break;
                }

                stdout.write("Description (optional):\n");
                var desc = stdin.readLineSync() ?? "";
                desc = desc.trim();

                var item = ScheduleItem(
                    DateTime.now().microsecondsSinceEpoch,
                    title,
                    desc
                );

                list.add(item);
                message = "New item was successfully added\n";
            } break;

            case "delete":
            case "del":
            case "d": {
                if (list.isEmpty) {
                    message = "The list is empty. You can add items using `add` command\n";
                    break;
                }

                isDeleting = true;
                String input = "";
                int? index = null;

                while (index == null) {
                    draw();
                    stdout.write("Select item to delete (empty to cancel): ");

                    input = stdin.readLineSync() ?? "";
                    input = input.trim().toLowerCase();
                    if (input.isEmpty) break;

                    index = int.tryParse(input);
                    if (index == null) {
                        message = "Invalid input: `$input`\n";
                    } else if (index > list.length) {
                        message = "Item `[$index]` is not valid\n";
                        index = null;
                    }
                }

                if (input.isEmpty) {
                    isDeleting = false;
                    message = "Canceled\n";
                    break;
                }

                index = index! - 1;
                list.removeAt(index);

                isDeleting = false;
                message = "Item successfully deleted\n";
            } break;

            case "q":
            case "quit":
            case "exit":
                exit = true;
                UI.ClearScreen();
                break;

            default: message = "Unknown command: `${cmd.name}`\n";
        }
    }
}

class Command {
    String name = "";
    List<String> args = [];

    Command(String cmd) {
        cmd = cmd.trim().toLowerCase();

        if (cmd.isEmpty) return;

        args = cmd.split(" ")
            .where((x) => x.isNotEmpty)
            .toList();

        name = args.first;
        args.removeAt(0);
    }
}

class ScheduleItem {
    int id;
    String title;
    String desc;

    ScheduleItem(this.id, this.title, this.desc);
}

class UI {
    UI.ClearScreen() {
        stdout.write(Process.runSync("clear", []).stdout);
    }

    UI.Text(String text) {
        stdout.write(text);
    }

    UI.EmptyLine() {
        UI.Text("\n");
    }

    UI.ListView(int listLen, Function(int i) builder) {
        for (int i = 0; i < listLen; i++) {
            builder(i);
        }
    }
}
