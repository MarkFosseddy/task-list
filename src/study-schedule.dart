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

        Text("┏━━ Study Schedule ━━━\n");
        EmptyLine();

        Text("  Morning:\n");
        ListView(morning.items.length, (int i) {
            var item = morning.items.elementAt(i);
            Text("    [${item.completed ? "X" : " "}] ${item.text}\n");
        });

        EmptyLine();

        Text("  Afternoon:\n");
        ListView(afternoon.items.length, (int i) {
            var item = afternoon.items.elementAt(i);
            Text("    [${item.completed ? "X" : " "}] ${item.text}\n");
        });

        EmptyLine();

        Text("  Evening:\n");
        ListView(evening.items.length, (int i) {
            var item = evening.items.elementAt(i);
            Text("    [${item.completed ? "X" : " "}] ${item.text}\n");
        });

        EmptyLine();
        EmptyLine();

        if (message.isNotEmpty) {
            Text(message);
        } else {
            EmptyLine();
        }

        Text(">> ");

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

            default: message = "Unknown command: `$cmd`\n";
        }
    }
}

// @TODO: load items from file
class Schedule {
    List<ScheduleItem> items;

    Schedule(this.items);
}

class ScheduleItem {
    String text;
    bool completed;

    ScheduleItem(this.text, this.completed);
}

class Text {
    Text(String text) {
        stdout.write(text);
    }
}

class EmptyLine {
    EmptyLine() {
        Text("\n");
    }
}

class ClearScreen {
    ClearScreen() {
        stdout.write(Process.runSync("clear", []).stdout);
    }
}

class ListView {
    ListView(int listLen, Function(int i) builder) {
        for (int i = 0; i < listLen; i++) {
            builder(i);
        }
    }
}
