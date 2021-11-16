import "dart:io";

void main() {
    bool exit = false;
    String cmd = "";
    String message = "";
    var schedule = Schedule();

    while (!exit) {
        // draw
        clearScreen();
        print("┏━━ Study Schedule ━━━");
        print("");
        schedule.draw();
        print("");
        print("");

        if (message.isNotEmpty) {
            print(message);
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
                clearScreen();
                break;

            default: message = "Unknown command: `$cmd`";
        }
    }
}

class Schedule {
    List<ScheduleItem> morning = [ScheduleItem("Hello, World!", true)];
    List<ScheduleItem> afternoon = [ScheduleItem("How Are You", false)];
    List<ScheduleItem> evening = [ScheduleItem("This is Me", false)];

    void draw() {
        if (this.morning.isNotEmpty) {
            drawLabel("Morning", this.morning);
        }

        if (this.afternoon.isNotEmpty) {
            drawLabel("Afternoon", this.afternoon);
        }

        if (this.evening.isNotEmpty) {
            drawLabel("Evening", this.evening);
        }
    }

    void drawLabel(String label, List<ScheduleItem> items) {
        print("  $label:");

        items.forEach((item) {
            print("    [${item.completed ? "X" : " "}] ${item.text}");
        });

        print("");
    }
}

class ScheduleItem {
    String text;
    bool completed;

    ScheduleItem(this.text, this.completed);
}

void clearScreen() {
    print(Process.runSync("clear", []).stdout);
}
