import "dart:io";

void main() {
    Program().run();
}

class Program {
    bool exit = false;
    Command cmd = Command("");
    String message = "";

    List<ScheduleItem> morning = [ScheduleItem(0, "Hello, World!", "")];
    List<ScheduleItem> afternoon = [ScheduleItem(0, "How Are You", "Optional description text")];
    List<ScheduleItem> evening = [ScheduleItem(0, "This is Me", "")];

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

        UI.Text("  Morning:\n");
        UI.ListView(morning.length, (int i) {
            var item = morning.elementAt(i);
            UI.Text("    ${item.title}\n");
            if (item.desc.isNotEmpty) {
                UI.Text("      ${item.desc}\n");
                UI.EmptyLine();
            } else {
                UI.EmptyLine();
            }
        });

        UI.EmptyLine();

        UI.Text("  Afternoon:\n");
        UI.ListView(afternoon.length, (int i) {
            var item = afternoon.elementAt(i);
            UI.Text("    ${item.title}\n");
            if (item.desc.isNotEmpty) {
                UI.Text("      ${item.desc}\n");
                UI.EmptyLine();
            } else {
                UI.EmptyLine();
            }
        });

        UI.EmptyLine();

        UI.Text("  Evening:\n");
        UI.ListView(evening.length, (int i) {
            var item = evening.elementAt(i);
            UI.Text("    ${item.title}\n");
            if (item.desc.isNotEmpty) {
                UI.Text("      ${item.desc}\n");
                UI.EmptyLine();
            } else {
                UI.EmptyLine();
            }
        });

        UI.EmptyLine();
        UI.EmptyLine();

        UI.Text("cmd-name: ${cmd.name}\n");
        UI.Text("cmd-args: ${cmd.args}\n");
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
            case "add": {
                if (cmd.args.isEmpty) {
                    message = "Not enough arguments for `add` command\n";
                    break;
                }

                if (!["m", "a", "e"].contains(cmd.args.first)) {
                    message = "Unknown argument `${cmd.args.first}` for `add` command\n";
                    break;
                }

                String title = "";
                String desc = "";

                stdout.write("Title: ");
                title = stdin.readLineSync() ?? "";
                title = title.trim();

                if (title.isEmpty) break;

                stdout.write("Description (optional):\n");
                desc = stdin.readLineSync() ?? "";
                desc = desc.trim();

                var item = ScheduleItem(
                    DateTime.now().microsecondsSinceEpoch,
                    title,
                    desc
                );

                switch (cmd.args.first) {
                    case "m":
                        morning.add(item);
                        break;

                    case "a":
                        afternoon.add(item);
                        break;

                    case "e":
                        evening.add(item);
                        break;
                }
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

        List<String> parts = cmd.split(" ")
            .where((x) => x.isNotEmpty)
            .toList();

        this.name = parts.first;

        parts.removeAt(0);
        this.args = parts;
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
