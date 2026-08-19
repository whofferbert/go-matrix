#!/usr/bin/env perl
# William Hofferbert

use strict;
use warnings;
use feature qw(say);
use Data::Dumper;

# turn off screen saver and blanking
`xset s off`;
`xset s noblank`;

# initiate lock, backgrounded
system("xtrlock &");

# get xtrlock pid
my $xtr_pid = `ps -ef | grep "[x]trlock" | awk '{print \$2}'`;
chomp $xtr_pid;

# bottom-right corner x:5279 y:2559
`xdotool mousemove 5279 2559`;

# the thing to run
my $matrix_cmd = "/home/whofferbert/etc/whoff/projects/matrix/matrix";

# call terms, fullscreen'd

# geometry notes
# right 142x127+3840+0 
# lower center 190x59+1920+1200
# upper center 190x59+1920+0
# lower left 190x59+0+1200
# upper left 190x59+0+0

my %geom_lines = (
  '1920x1200' => ["190", "59"],
  '1440x2560' => ["142", "127"],

);

say Dumper(\%geom_lines);

# get original pid for mate-terminal
my $original_pid = `ps -ef | grep "[m]ate-terminal" | awk '{print \$2}'`;
chomp $original_pid;

sub get_screen_geoms {
  my $out = `xrandr --listactivemonitors`;
  my @xys;
  for my $line (split(/\n/, $out)) {
    my @parts = split(/[\/+x\s]/, $line);
    say Dumper(\@parts);
    if (scalar(@parts) > 8) {
      my %h;
      $h{res_x} = $parts[4];
      $h{res_y} = $parts[6];
      $h{off_x} = $parts[8];
      $h{off_y} = $parts[9];
      push @xys, \%h; 
    }
  }
  return @xys;
}

&get_screen_geoms;

for my $g (&get_screen_geoms) {
  my $xy = $g->{res_x} . "x" . $g->{res_y};
  my $rows_cols = $geom_lines{$xy};
  my $geom = $rows_cols->[0] . "x" . $rows_cols->[1] . "+" . $g->{off_x} . "+" . $g->{off_y};
  my $cmd = "";
  $cmd .= "mate-terminal --hide-menubar --full-screen ";
  $cmd .= "--geometry $geom --profile matrix ";
  $cmd .= "--disable-factory -e $matrix_cmd &";
  system($cmd);
  #say $cmd;
}

my $running = 1;

while ($running) {
  if (! -e "/proc/$xtr_pid") {
    $running = 0;
  }
  # TODO sleep 0.25
  sleep 1;
}

# once done, collect pids for mate-terminals called to run 'matrix' and kill em
`ps -ef | grep -P "[m]ate-terminal.*matrix" | awk '{print \$2}' | grep -Pv "$original_pid" | xargs kill`;

# set screensaver and blanking back on
`xset s on`;
`xset s blank`;
