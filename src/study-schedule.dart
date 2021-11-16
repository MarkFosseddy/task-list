import "dart:io";

void main() {
    bool exit = false;
    String cmd = "";
    String message = "";

    var morning = Schedule([ScheduleItem("Hello, World!", true)]);
    var afternoon = Schedule([ScheduleItem("How Are You", false)]);
    var evening = Schedule([ScheduleItem("This is Me", false)]);

    while (!exit) {
        // draw
        ClearScreen();

        Text("┏━━ Study Schedule ━━━");
        NewLine();

        Text("  Morning:");
        ListView(morning.items, (item) {
            Text("    [${item.completed ? "X" : " "}] ${item.text}");
        });
        NewLine();

        Text("  Afternoon:");
        ListView(afternoon.items, (item) {
            Text("    [${item.completed ? "X" : " "}] ${item.text}");
        });
        NewLine();

        Text("  Evening:");
        ListView(evening.items, (item) {
            Text("    [${item.completed ? "X" : " "}] ${item.text}");
        });

        NewLine();
        NewLine();

        if (message.isNotEmpty) {
            Text(message);
        }

        // update
        message = "";

        cmd = stdin.readLineSync() ?? "";
        cmd = cmd.trim().toLowerCase();
        if (cmd.isEmpty) continue;

        switch (cmd) {
            case "add":
                break;

            case "q":
                exit = true;
                ClearScreen();
                break;

            default: message = "Unknown command: `$cmd`";
        }
    }
}

class Text {
    Text(String text) {
        print(text);
    }
}

class NewLine {
    NewLine() {
        Text("");
    }
}

class ClearScreen {
    ClearScreen() {
        print(Process.runSync("clear", []).stdout);
    }
}

class ListView {
    ListView(List<dynamic> data, Function(dynamic v) builder) {
        data.forEach(builder);
    }
}

class Schedule {
    List<ScheduleItem> items;

    Schedule(this.items);
}

class ScheduleItem {
    String text;
    bool completed;

    ScheduleItem(this.text, this.completed);
}
